#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { ReportGeneratorStack } from '../lib/cdk-stack';

const app = new cdk.App();
new ReportGeneratorStack(app, 'ReportGeneratorStack', {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
});
