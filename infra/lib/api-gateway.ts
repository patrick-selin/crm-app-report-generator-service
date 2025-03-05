import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import * as apigateway from "aws-cdk-lib/aws-apigateway";
import { ReportLambdas } from "./lambdas";

interface ApiGatewayProps {
  reportLambdas: ReportLambdas;
}

export class ApiGateway extends Construct {
  constructor(scope: Construct, id: string, props: ApiGatewayProps) {
    super(scope, id);

    const api = new apigateway.RestApi(this, "ReportApi", {
      restApiName: "Report Generator Service",
    });

    // Base resource: /reports
    const reportsResource = api.root.addResource("reports");

    // Create sub-resource: /reports/new
    const newResource = reportsResource.addResource("new");
    newResource.addMethod(
      "POST",
      new apigateway.LambdaIntegration(props.reportLambdas.reportLambda)
    );

    new cdk.CfnOutput(this, 'ApiEndpoint', {
      value: api.url, 
      description: 'The base URL of the API Gateway',
    });
  }
}