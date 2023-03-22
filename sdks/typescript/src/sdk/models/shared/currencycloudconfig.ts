/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class CurrencyCloudConfig extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "apiKey" })
  apiKey: string;

  /**
   * The endpoint to use for the API. Defaults to https://devapi.currencycloud.com
   */
  @SpeakeasyMetadata()
  @Expose({ name: "endpoint" })
  endpoint?: string;

  /**
   * Username of the API Key holder
   */
  @SpeakeasyMetadata()
  @Expose({ name: "loginID" })
  loginID: string;

  /**
   * The frequency at which the connector will fetch transactions
   */
  @SpeakeasyMetadata()
  @Expose({ name: "pollingPeriod" })
  pollingPeriod?: string;
}