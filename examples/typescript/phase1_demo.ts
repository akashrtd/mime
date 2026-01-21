/**
 * Phase 1 Feature Demo - Markdown, Metadata, Links, Cookies
 * 
 * This example demonstrates the new Phase 1 features added to MIME
 */

import { MIME } from '@mime-browser/sdk';

async function main() {
    console.log('🚀 MIME Phase 1 Feature Demo\n');
    
    const browser = await MIME.connect();

    try {
        // Navigate to a page
        console.log('📍 Navigating to Hacker News...');
        await browser.navigate('https://news.ycombinator.com');
        await browser.waitFor('text=Hacker News');

        // Test metadata extraction
        console.log('\n📋 Page Metadata:');
        const meta = await browser.metadata();
        console.log(`  Title: ${meta.title}`);
        console.log(`  URL: ${meta.url}`);
        console.log(`  Description: ${meta.description || '(none)'}`);

        // Test markdown extraction
        console.log('\n📝 Markdown Content (first 500 chars):');
        const { markdown } = await browser.markdown();
        console.log(markdown.slice(0, 500) + '...');

        // Test links extraction
        console.log('\n🔗 Links (first 5):');
        const { links, count } = await browser.links();
        console.log(`  Total links found: ${count}`);
        links.slice(0, 5).forEach((link, i) => {
            console.log(`  ${i + 1}. ${link.text.slice(0, 50)} -> ${link.url.slice(0, 60)}`);
        });

        // Test cookies
        console.log('\n🍪 Cookies:');
        const { cookies } = await browser.getCookies();
        console.log(`  Found ${cookies.length} cookies`);
        cookies.slice(0, 3).forEach(c => {
            console.log(`    - ${c.name}: ${c.value.slice(0, 20)}...`);
        });

        console.log('\n✅ All Phase 1 features working!');

    } finally {
        await browser.close();
    }
}

main().catch(console.error);
