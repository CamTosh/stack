/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class MangoPayConfig extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "apiKey" })
  apiKey: string;

  @SpeakeasyMetadata()
  @Expose({ name: "clientID" })
  clientID: string;

  @SpeakeasyMetadata()
  @Expose({ name: "endpoint" })
  endpoint: string;
}