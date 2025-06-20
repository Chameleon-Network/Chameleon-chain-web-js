import { addressAsObject, KeySet } from "../../../../common/keySet";
import {
  BurningPRVERC20RequestMeta,
  BurningPRVBEP20RequestMeta,
  encryptMessageOutCoin,
  getBurningAddress,
} from "../../../../core";
import Validator from "../../../../utils/validator";
import bn from "bn.js";
import { TX_TYPE } from "../../account.constants";
import { PrivacyVersion, PRVIDSTR } from "../../../../core/constants";

// createAndSendBurningPegPRVRequestTx create and send tx burning native PRV token to receive pegPRV
// remoteAddress (string) is an ERC20/BEP20 address which users want to receive pegPRV (without 0x)

/** createAndSendBurningPegPRVRequestTx
 * @param {{paymentAddressStr: string, amount: number}} prvPayments
 * @param {number} fee
 * @param {string} info
 * @param {string} tokenID
 * @param {object} metadata
 * @param {boolean} burningType
 * @param {boolean} isEncryptMessage
 * @param {number} burnAmount
 * @param {string} remoteAddress
 *
 */
async function createAndSendBurningPegPRVRequestTx({
  transfer: { fee, tokenID = PRVIDSTR, info = "", prvPayments = [] },
  extra: {
    burningType,
    burnAmount,
    remoteAddress,
    txHashHandler = null,
    version = PrivacyVersion.ver2,
    isEncryptMessage = true,
  } = {},
}) {
  try {
    new Validator("createAndSendBurningPegPRVRequestTx-fee", fee).required().amount();
    new Validator("createAndSendBurningPegPRVRequestTx-info", info).string();
    new Validator("createAndSendBurningPegPRVRequestTx-tokenID", tokenID)
      .required()
      .inList([PRVIDSTR]);
    new Validator(
      "createAndSendBurningPegPRVRequestTx-burningType",
      burningType
    ).inList([BurningPRVERC20RequestMeta, BurningPRVBEP20RequestMeta]);
    new Validator("createAndSendBurningPegPRVRequestTx-burnAmount", burnAmount)
      .required()
      .amount();
    new Validator("createAndSendBurningPegPRVRequestTx-remoteAddress", remoteAddress)
      .required()
      .string();
    new Validator("createAndSendBurningPegPRVRequestTx-prvPayments", prvPayments)
      .required()
      .array();
    new Validator("createAndSendBurningPegPRVRequestTx-version", version)
      .required()
      .number();
    const burningTokenID = tokenID;
    if (remoteAddress.startsWith("0x")) {
      remoteAddress = remoteAddress.slice(2);
    }
    await this.updateProgressTx(10, "Encrypting Message");
    const burningAddress = await getBurningAddress(this.rpc);
    prvPayments = prvPayments.map((payment) => ({
      PaymentAddress: payment?.paymentAddress,
      Amount: new bn(payment?.amount).toString(),
      Message: "",
    }));
    prvPayments.push({
      PaymentAddress: burningAddress,
      Amount: new bn(burnAmount).toString(),
      Message: "",
    });
    console.log("prvPayments: ", prvPayments);
    const isEncodeOnly = !isEncryptMessage;
    prvPayments = await encryptMessageOutCoin(prvPayments, isEncodeOnly);

    // use an empty payment address
    let emptyKeySet = new KeySet();
    await emptyKeySet.importFromPrivateKey(new Uint8Array(32));
    let addrForMd = addressAsObject(emptyKeySet.PaymentAddress);
    await this.updateProgressTx(15, "Generating Metadata");
    // prepare meta data for tx
    const burningReqMetadata = {
      BurnerAddress: addrForMd,
      BurningAmount: burnAmount,
      TokenID: burningTokenID,
      RemoteAddress: remoteAddress,
      Type: burningType,
    };
    console.log("burningReqMetadata", burningReqMetadata);
    const tx = await this.transact({
      transfer: {
        prvPayments,
        fee,
        info,
        tokenID: burningTokenID,
      },
      extra: {
        metadata: burningReqMetadata,
        txType: TX_TYPE.BURN,
        txHashHandler,
        version,
      },
    });
    await this.updateProgressTx(100, "Completed");
    return tx;
  } catch (error) {
    throw error;
  }
}
export default {
  createAndSendBurningPegPRVRequestTx,
};
