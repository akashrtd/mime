/**
 * Phase 3 Feature Demo - Crawl and Map
 * 
 * This example demonstrates the Multi-Page tools: crawl and map
 */

import { MIME } from '@mime-browser/sdk';

async function main() {
    console.log('🕷️ MIME Phase 3 Feature Demo - Multi-Page Tools\n');
    
    const browser = await MIME.connect();

    try {
        const startUrl = 'https://news.ycombinator.com';

        // Test Map tool
        console.log(`🗺️ Mapping site structure from: ${startUrl}`);
        const mapResult = await browser.map(startUrl);
        
        console.log(`  Found ${mapResult.total} URLs`);
        console.log('  First 10 URLs:');
        mapResult.urls.slice(0, 10).forEach(u => console.log(`    - ${u}`));

        // Test Crawl tool
        console.log(`\n🕸️ Crawling pages (shallow, just top 3)...`);
        const crawlResult = await browser.crawl(startUrl, {
            max_pages: 3,
            max_depth: 1,
            same_domain: true
        });
        
        console.log(`  Crawled ${crawlResult.total} pages`);
        
        crawlResult.pages.forEach((p, i) => {
            console.log(`\n  Page ${i+1}:`);
            console.log(`    URL: ${p.url}`);
            console.log(`    Title: ${p.title}`);
            console.log(`    Status: ${p.status}`);
            console.log(`    Markdown length: ${p.markdown?.length || 0} chars`);
        });

        console.log('\n✅ Phase 3 Multi-Page tools working!');

    } finally {
        await browser.close();
    }
}

main().catch(console.error);
