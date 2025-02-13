// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// Do not modify this file. This file is machine generated, and any changes to it will be overwritten.
package software.amazon.cryptography.keystore.model;

import java.util.Objects;

/**
 * Outputs for Key Store DynamoDB table creation.
 */
public class CreateKeyStoreOutput {

  /**
   * The ARN of the DynamoDB table that backs this Key Store.
   */
  private final String tableArn;

  protected CreateKeyStoreOutput(BuilderImpl builder) {
    this.tableArn = builder.tableArn();
  }

  /**
   * @return The ARN of the DynamoDB table that backs this Key Store.
   */
  public String tableArn() {
    return this.tableArn;
  }

  public Builder toBuilder() {
    return new BuilderImpl(this);
  }

  public static Builder builder() {
    return new BuilderImpl();
  }

  public interface Builder {
    /**
     * @param tableArn The ARN of the DynamoDB table that backs this Key Store.
     */
    Builder tableArn(String tableArn);

    /**
     * @return The ARN of the DynamoDB table that backs this Key Store.
     */
    String tableArn();

    CreateKeyStoreOutput build();
  }

  static class BuilderImpl implements Builder {

    protected String tableArn;

    protected BuilderImpl() {}

    protected BuilderImpl(CreateKeyStoreOutput model) {
      this.tableArn = model.tableArn();
    }

    public Builder tableArn(String tableArn) {
      this.tableArn = tableArn;
      return this;
    }

    public String tableArn() {
      return this.tableArn;
    }

    public CreateKeyStoreOutput build() {
      if (Objects.isNull(this.tableArn())) {
        throw new IllegalArgumentException(
          "Missing value for required field `tableArn`"
        );
      }
      if (Objects.nonNull(this.tableArn()) && this.tableArn().length() < 1) {
        throw new IllegalArgumentException(
          "The size of `tableArn` must be greater than or equal to 1"
        );
      }
      if (Objects.nonNull(this.tableArn()) && this.tableArn().length() > 1024) {
        throw new IllegalArgumentException(
          "The size of `tableArn` must be less than or equal to 1024"
        );
      }
      return new CreateKeyStoreOutput(this);
    }
  }
}
