const puppeteer = require('puppeteer');

const ITERATIONS = 5;
const TEST_URL = 'https://example.com';

async function benchmark() {
    console.log('=== Puppeteer Benchmark ===\n');
    
    const results = {
        startup: [],
        navigation: [],
        extraction: [],
        screenshot: [],
        total: []
    };

    for (let i = 0; i < ITERATIONS; i++) {
        console.log(`Run ${i + 1}/${ITERATIONS}`);
        
        const runStart = Date.now();
        
        // Startup time
        const startupStart = Date.now();
        const browser = await puppeteer.launch({ headless: true });
        const page = await browser.newPage();
        results.startup.push(Date.now() - startupStart);

        // Navigation time
        const navStart = Date.now();
        await page.goto(TEST_URL, { waitUntil: 'networkidle0' });
        results.navigation.push(Date.now() - navStart);

        // Extraction time
        const extractStart = Date.now();
        const title = await page.$eval('h1', el => el.textContent);
        const html = await page.content();
        const links = await page.$$eval('a', els => els.map(e => e.href));
        results.extraction.push(Date.now() - extractStart);

        // Screenshot time
        const screenshotStart = Date.now();
        await page.screenshot({ encoding: 'base64' });
        results.screenshot.push(Date.now() - screenshotStart);

        await browser.close();
        
        results.total.push(Date.now() - runStart);
    }

    // Calculate averages
    console.log('\n=== Results (average of ' + ITERATIONS + ' runs) ===');
    console.log(`Startup:     ${avg(results.startup).toFixed(0)} ms`);
    console.log(`Navigation:  ${avg(results.navigation).toFixed(0)} ms`);
    console.log(`Extraction:  ${avg(results.extraction).toFixed(0)} ms`);
    console.log(`Screenshot:  ${avg(results.screenshot).toFixed(0)} ms`);
    console.log(`Total:       ${avg(results.total).toFixed(0)} ms`);
    
    // Output JSON for comparison
    console.log('\nJSON:');
    console.log(JSON.stringify({
        tool: 'puppeteer',
        iterations: ITERATIONS,
        averages: {
            startup: avg(results.startup),
            navigation: avg(results.navigation),
            extraction: avg(results.extraction),
            screenshot: avg(results.screenshot),
            total: avg(results.total)
        },
        raw: results
    }));
}

function avg(arr) {
    return arr.reduce((a, b) => a + b, 0) / arr.length;
}

benchmark().catch(console.error);
