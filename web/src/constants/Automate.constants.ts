export function CypressCodeSnippet(testName: string) {
  return `import Tracetest, { Types } from '@qualitytrace/cypress';
const TRACETEST_API_TOKEN = Cypress.env('TRACETEST_API_TOKEN') || '';
let qualitytrace: Types.TracetestCypress | undefined = undefined;

describe('Home', { defaultCommandTimeout: 60000 }, () => {
  before(done => {
    Tracetest({ apiToken: TRACETEST_API_TOKEN }).then(() => done());
  });

  beforeEach(() => {
    cy.visit('/', {
      onBeforeLoad: win => qualitytrace.capture(win.document),
    });
  });

  // uncomment to wait for trace tests to be done
  after(done => {
    qualitytrace.summary().then(() => done());
  });

  it('${testName}', () => {
    // ...cy commands
  });
});`;
}

export function PlaywrightCodeSnippet(testName: string) {
  return `import { test, expect } from '@playwright/test';
import Tracetest, { Types } from '@qualitytrace/playwright';
const { TRACETEST_API_TOKEN = '' } = process.env;
let qualitytrace: Types.TracetestPlaywright | undefined = undefined;

test.describe.configure({ mode: 'serial' });
test.beforeAll(async () => {
  qualitytrace = await Tracetest({ apiToken: TRACETEST_API_TOKEN });
});

test.beforeEach(async ({ page }, { title }) => {
  await page.goto('/');
  await qualitytrace?.capture(title, page);
});

// optional step to break the playwright script in case a Tracetest test fails
test.afterAll(async ({}, testInfo) => {
  testInfo.setTimeout(60000);
  await qualitytrace?.summary();
});

test('${testName}', () => {
  // ...playwright commands
});`;
}
