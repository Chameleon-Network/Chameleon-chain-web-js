import Validator from '../../../../utils/validator';
import { LIMIT_DEFAULT } from './liquidity.constant';
import flatten from 'lodash/flatten';
import {
  MAP_REMOVE_LP_HISTORIES_SERVICE,
  MAP_STORAGE_REMOVE_LP_HISTORIES,
  REMOVE_LP_HISTORIES_TX_STR,
  TX_TYPE
} from '../../../Account';
import orderBy from 'lodash/orderBy';

async function getRemoveLPHistories({ offset = 0, limit = LIMIT_DEFAULT } = {}) {
  new Validator("getRemoveLPHistories-offset", offset).required().number();
  new Validator("getRemoveLPHistories-limit", limit).required().number();
  let [csvHistories, storageHistories] = await Promise.all([
    flatten(await this.getRemoveLPHistoriesApi()),
    await this.getStorageHistoriesRemoveLP(),
  ])
  const requestTxs = csvHistories.map(({ requestTx }) => requestTx);
  storageHistories = storageHistories.filter(({ requestTx, type }) =>
    (type === TX_TYPE.WITHDRAW_CONTRIBUTE && !requestTxs.includes(requestTx))
  );
  storageHistories = await Promise.all(storageHistories.map(async history => {
    let status = await this.account?.rpcTxService.apiGetTxStatus({ txId: history.requestTx });
    status = MAP_STORAGE_REMOVE_LP_HISTORIES[status];
    const statusStr = REMOVE_LP_HISTORIES_TX_STR[status];
    return {
      ...history,
      status,
      statusStr,
    }
  }));
  csvHistories = csvHistories.map(history => {
    const status = MAP_REMOVE_LP_HISTORIES_SERVICE[history.status];
    const statusStr = REMOVE_LP_HISTORIES_TX_STR[status];
    const requestTime = history.requestime;
    return {
      ...history,
      status,
      statusStr,
      requestTime,
    }
  })
  return orderBy(csvHistories.concat(storageHistories), 'requestTime', ['desc']);
}

export default ({
  getRemoveLPHistories,
})
