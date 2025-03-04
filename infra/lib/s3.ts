import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as s3 from "aws-cdk-lib/aws-s3";

export class S3Resources extends Construct {
  public readonly reportsBucket: s3.Bucket;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.reportsBucket = new s3.Bucket(this, "ReportsBucket", {
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      autoDeleteObjects: false,
    });
  }
}
