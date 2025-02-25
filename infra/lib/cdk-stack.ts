import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';

export class ReportGeneratorStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // DynamoDB Table
    // S3 Bucket
    // Lambda Function for Report Generation
    // Lambda Function for Report Status
    // Grant permissions
    // API Gateway


  }
}
