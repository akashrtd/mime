import { MIME } from '@mime-browser/sdk';

async function searchPetFood() {
    console.log('Connecting to MIME...');
    const browser = await MIME.connect();

    try {
        console.log('Navigating to DuckDuckGo (HTML)...');
        // Use HTML version for better bot compatibility
        await browser.navigate('https://html.duckduckgo.com/html/');

        console.log('Typing query...');
        await browser.waitFor('input[name="q"]');
        await browser.type('input[name="q"]', 'pet food');

        console.log('Submitting search...');
        // Submit the form directly using JavaScript (wrapped in function)
        await browser.execute(`() => { document.querySelector('form[action="/html/"]').submit(); return true; }`);

        console.log('Waiting for results...');
        await browser.waitFor('.result__title');

        console.log('Extracting results...');
        // Extract top 3 result titles
        const results = await browser.execute(`
            () => {
                return Array.from(document.querySelectorAll('.result__title')).slice(0, 3).map(el => el.textContent.trim());
            }
        `) as string[];

        console.log('\nTop 3 Results for "pet food":');
        results.forEach((title, i) => console.log(`${i+1}. ${title}`));

    } catch (e) {
        console.error('Search failed:', e);
    } finally {
        await browser.close();
    }
}

searchPetFood();
