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

    const reportResource = api.root.addResource("reports");
    reportResource.addMethod(
      "POST",
      new apigateway.LambdaIntegration(props.reportLambdas.reportLambda)
    );
  }
}
