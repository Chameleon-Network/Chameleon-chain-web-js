import { PrivacyVersion } from "../../../../core/constants";
import Validator from "../../../../utils/validator";

async function checkKeyImages(params) {
  try {
    const { version, listOutputsCoins, shardId, tokenID } = params;
    new Validator(`checkKeyImages-shardId`, shardId).required().number();
    new Validator(`checkKeyImages-version`, version).required().number();
    new Validator(`checkKeyImages-listOutputsCoins`, listOutputsCoins)
      .required()
      .array();
    new Validator(`checkKeyImages-tokenID`, tokenID).required().string();
    switch (version) {
      case PrivacyVersion.ver1:
        return this.checkKeyImageV1(params);
      case PrivacyVersion.ver2:
        return this.checkKeyImageV2(params);
      default:
        break;
    }
  } catch (error) {
    throw error;
  }
}

export default {
  checkKeyImages,
};
