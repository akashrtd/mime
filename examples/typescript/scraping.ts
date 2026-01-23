/**
 * Web Scraping with MIME TypeScript SDK
 * 
 * This example demonstrates scraping data from multiple pages
 */

import { MIME } from '@mime-browser/sdk';

interface Article {
    title: string;
    url: string;
}

async function scrapeHackerNews(): Promise<Article[]> {
    // Use absolute path to binary since it's not in PATH
    const mimePath = process.env.MIME_PATH || '/Users/akashrathod/Desktop/projects/mime/bin/mime';
    const browser = await MIME.connect({ binaryPath: mimePath });
    const articles: Article[] = [];

    try {
        await browser.navigate('https://news.ycombinator.com');

        // Wait for the logo text to ensure page load (Smart Selector demo)
        await browser.waitFor('text=Hacker News');

        // Scroll down to trigger any lazy loading (Demo)
        await browser.scroll(undefined, 0, 500);

        // Execute JavaScript to get all titles and links
        const result = await browser.execute(`
          () => {
            return Array.from(document.querySelectorAll('.titleline a')).slice(0, 10).map(a => ({
                title: a.textContent,
                url: a.href
            }));
          }
        `) as Article[];
        
        articles.push(...result);

        // articles.push(...result);

    } finally {
        await browser.close();
    }

    return articles;
}

async function main() {
    console.log('Scraping Hacker News top 10 articles...\n');

    const articles = await scrapeHackerNews();

    articles.forEach((article, i) => {
        console.log(`${i + 1}. ${article.title}`);
        console.log(`   ${article.url}\n`);
    });
}

main().catch(console.error);
