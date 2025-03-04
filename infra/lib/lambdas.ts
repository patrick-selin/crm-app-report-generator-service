import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as iam from "aws-cdk-lib/aws-iam";
import * as path from "path";
import { S3Resources } from "./s3";
import { DynamoDBResources } from "./dynamodb";

interface LambdaProps {
  s3Resources: S3Resources;
  dynamoDBResources: DynamoDBResources;
}

export class ReportLambdas extends Construct {
  public readonly reportLambda: lambda.Function;

  constructor(scope: Construct, id: string, props: LambdaProps) {
    super(scope, id);

    // Lambda Role & Policies
    const lambdaRole = new iam.Role(this, "LambdaExecutionRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
    });

    lambdaRole.addManagedPolicy(
      iam.ManagedPolicy.fromAwsManagedPolicyName(
        "service-role/AWSLambdaBasicExecutionRole"
      )
    );

    props.s3Resources.reportsBucket.grantReadWrite(lambdaRole);
    props.dynamoDBResources.reportMetadataTable.grantReadWriteData(lambdaRole);

    this.reportLambda = new lambda.Function(this, "ReportLambda", {
      runtime: lambda.Runtime.PROVIDED_AL2,
      handler: "bootstrap",
      code: lambda.Code.fromAsset(path.join(__dirname, "../../lambda/report-generator"), {
        bundling: {
          image: lambda.Runtime.PROVIDED_AL2.bundlingImage,
          command: [
            "bash",
            "-c",
            "GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /asset-output/bootstrap main.go",
          ],
          volumes: [
            {
              hostPath: "/tmp/cdk-go-build",
              containerPath: "/go/pkg/mod",
            },
          ],
          environment: {
            GOCACHE: "/go/pkg/mod",
          },
        },
      }),
      environment: {
        BUCKET_NAME: props.s3Resources.reportsBucket.bucketName,
        TABLE_NAME: props.dynamoDBResources.reportMetadataTable.tableName,
      },
      role: lambdaRole,
    });
  }
}
