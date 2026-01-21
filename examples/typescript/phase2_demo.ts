/**
 * Phase 2 Feature Demo - Observe and Act
 * 
 * This example demonstrates the AI agent tools: observe and act
 */

import { MIME } from '@mime-browser/sdk';

async function main() {
    console.log('🤖 MIME Phase 2 Feature Demo - AI Agent Tools\n');
    
    const browser = await MIME.connect();

    try {
        // Navigate to a page with forms
        console.log('📍 Navigating to Hacker News login page...');
        await browser.navigate('https://news.ycombinator.com/login');
        await browser.waitFor('text=Login');

        // Test observe tool
        console.log('\n👁️ Observing page structure...');
        const obs = await browser.observe();
        
        console.log(`  URL: ${obs.url}`);
        console.log(`  Title: ${obs.title}`);
        console.log(`  Forms found: ${obs.forms.length}`);
        
        if (obs.forms.length > 0) {
            const form = obs.forms[0];
            console.log(`  First form fields:`);
            form.fields.forEach(f => {
                console.log(`    - ${f.name || f.placeholder || f.type} (${f.selector})`);
            });
            if (form.submit) {
                console.log(`  Submit button: "${form.submit.text}"`);
            }
        }

        console.log(`\n  Clickable elements: ${obs.clickable.length}`);
        obs.clickable.slice(0, 5).forEach(c => {
            console.log(`    - "${c.text}" (${c.type})`);
        });

        console.log(`\n  Content headings: ${obs.content.headings.join(', ') || '(none)'}`);

        // Test act tool with natural language
        console.log('\n🎯 Testing natural language actions...');
        
        // Try to click something
        const result = await browser.act('click forgot password');
        console.log(`  Action: ${result.action}`);
        console.log(`  Success: ${result.success}`);
        console.log(`  Target: ${result.target}`);
        if (result.message) {
            console.log(`  Message: ${result.message}`);
        }

        console.log('\n✅ Phase 2 AI Agent tools working!');

    } finally {
        await browser.close();
    }
}

main().catch(console.error);
