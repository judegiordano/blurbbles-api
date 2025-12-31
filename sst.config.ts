/// <reference path='./.sst/platform/config.d.ts' />

const domain = 'blurbbles.com'

export default $config({
  app(input) {
    return {
      name: 'blurbbles-api',
      removal: 'remove',
      home: 'aws',
    };
  },
  async run() {
    const { stage } = $app
    const environment = {
      STAGE: stage,
    }

    const api = new sst.aws.Function('blurbbles-api', {
      runtime: 'go',
      architecture: 'arm64',
      memory: '1 GB',
      timeout: '10 minutes',
      url: {
        cors: {
          allowMethods: ['*'],
          allowOrigins: ['*'],
          allowCredentials: true,
          maxAge: '1 day',

        }
      },
      handler: './lambdas/api/main.go',
      logging: {
        retention: '1 week',
        format: 'json'
      },
      environment
    })

    const router = new sst.aws.Router('blurbbles-api-router', {
      invalidation: false,
      routes: { '/*': api.url },
      domain: {
        name: `api.${domain}`,
        redirects: [`www.api.${domain}`]
      }
    })

    return {
      function: api.url,
      api: router.url
    }
  },
});
