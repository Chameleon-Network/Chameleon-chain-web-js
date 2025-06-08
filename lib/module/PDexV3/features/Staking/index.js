import data from './staking.data';
import histories from './staking.histories';
import transactions from './staking.transaction';
import storage from './staking.storage';

export default {
  ...data,
  ...histories,
  ...transactions,
  ...storage,
}
