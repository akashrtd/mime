/**
 * Basic MIME TypeScript SDK Example
 * 
 * This example demonstrates basic browser automation:
 * - Navigating to a page
 * - Extracting content
 * - Taking screenshots
 */

import { MIME } from '@mime-browser/sdk';
import * as fs from 'fs';

async function main() {
    console.log('Connecting to MIME...');

    // Connect to MIME server
    const browser = await MIME.connect();

    try {
        // Navigate to Hacker News
        console.log('Navigating to Hacker News...');
        await browser.navigate('https://news.ycombinator.com');

        // Extract the top story title
        const topStory = await browser.extract('.titleline a');
        console.log(`Top story: ${topStory}`);

        // Get current URL
        const url = await browser.url();
        console.log(`Current URL: ${url}`);

        // Take a screenshot
        console.log('Taking screenshot...');
        const { screenshot } = await browser.screenshot();

        // Save screenshot to file
        const buffer = Buffer.from(screenshot, 'base64');
        fs.writeFileSync('hackernews.png', buffer);
        console.log('Screenshot saved to hackernews.png');

    } finally {
        // Always close the browser
        await browser.close();
        console.log('Done!');
    }
}

main().catch(console.error);
