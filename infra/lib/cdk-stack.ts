import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { S3Resources } from "./s3";
import { DynamoDBResources } from "./dynamodb";
import { ReportLambdas } from "./lambdas";
import { ApiGateway } from "./api-gateway";

export class ReportGeneratorStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const s3Resources = new S3Resources(this, "S3Resources");
    const dynamoDBResources = new DynamoDBResources(this, "DynamoDBResources");

    const reportLambdas = new ReportLambdas(this, "ReportLambdas", {
      s3Resources,
      dynamoDBResources,
    });

    new ApiGateway(this, "ApiGateway", {
      reportLambdas,
    });
  }
}
