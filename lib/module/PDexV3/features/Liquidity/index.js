import histories from './liquidity.histories';
import historiesService from './liquidity.historiesService';
import storage from './liquidity.storage';
import transaction from './liquidity.transaction';
import historiesContribute from './liquidity.historiesContribute';
import historiesRemovePool from './liquidity.historiesRemovePool';
import historiesWithdrawReward from './liquidity.historiesWithdrawReward';

export default {
  ...histories,
  ...historiesService,
  ...storage,
  ...transaction,
  ...historiesContribute,
  ...historiesRemovePool,
  ...historiesWithdrawReward,
};
