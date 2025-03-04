import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as dynamodb from "aws-cdk-lib/aws-dynamodb";

export class DynamoDBResources extends Construct {
  public readonly reportMetadataTable: dynamodb.Table;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.reportMetadataTable = new dynamodb.Table(this, "ReportMetadataTable", {
      partitionKey: { name: "ReportID", type: dynamodb.AttributeType.STRING },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    });
  }
}
