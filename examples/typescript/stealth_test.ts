import { MIME } from '@mime-browser/sdk';

async function main() {
    console.log('🕵️  MIME Stealth Mode Test\n');
    
    // Pass custom args to enable stealth
    const browser = await MIME.connect({
        args: ['serve', '--stealth']
    });

    try {
        console.log('📍 Checking navigator.webdriver...');
        const isWebDriver = await browser.execute('() => navigator.webdriver');
        console.log(`  navigator.webdriver = ${isWebDriver}`);
        
        if (isWebDriver === false || isWebDriver === undefined) {
            console.log('✅ Stealth Mode working! (webdriver flag hidden)');
        } else {
            console.log('❌ Stealth Mode FAILED! (webdriver flag visible)');
            process.exit(1);
        }

    } finally {
        await browser.close();
    }
}

main().catch(console.error);
