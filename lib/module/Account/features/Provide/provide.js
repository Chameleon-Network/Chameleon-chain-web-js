import Validator from "../../../../utils/validator";
import { wasm } from "../../../../wasm";

async function signPoolWithdraw({ amount }) {
  new Validator("amount", amount).required();
  const privateKey = this.getPrivateKey();
  const paymentAddress = this.getPaymentAddress();
  const args = {
    data: {
      privateKey,
      paymentAddress,
      amount,
    },
  };
  return wasm.signPoolWithdraw(JSON.stringify(args));
}

export default {
  signPoolWithdraw,
};
