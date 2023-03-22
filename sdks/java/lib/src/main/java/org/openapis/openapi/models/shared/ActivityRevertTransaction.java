/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityRevertTransaction {
    @JsonProperty("id")
    public String id;

    public ActivityRevertTransaction withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonProperty("ledger")
    public String ledger;

    public ActivityRevertTransaction withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    public ActivityRevertTransaction(@JsonProperty("id") String id, @JsonProperty("ledger") String ledger) {
        this.id = id;
        this.ledger = ledger;
  }
}