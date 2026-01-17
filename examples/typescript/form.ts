/**
 * Form Automation with MIME TypeScript SDK
 * 
 * This example demonstrates filling out and submitting forms
 */

import { MIME } from '@mime-browser/sdk';

async function loginExample() {
    const browser = await MIME.connect();

    try {
        // Navigate to a login page (example.com doesn't have one, this is illustrative)
        await browser.navigate('https://example.com');

        // Wait for page to load
        await browser.waitFor('h1');

        // Example of form filling (uncomment for real forms):
        // await browser.type('#username', 'myuser');
        // await browser.type('#password', 'mypassword');
        // await browser.click('#login-button');

        // Get the page content
        const title = await browser.extract('h1');
        console.log(`Page title: ${title}`);

        // Get full HTML
        const { html, url } = await browser.html();
        console.log(`URL: ${url}`);
        console.log(`HTML length: ${html.length} characters`);

    } finally {
        await browser.close();
    }
}

async function main() {
    console.log('Form Automation Example\n');
    await loginExample();
    console.log('\nDone!');
}

main().catch(console.error);
