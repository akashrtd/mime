/**
 * MIME Feature Validation Test
 * Tests all 19 MCP tools to ensure they work correctly
 */

import { MIME } from "../../sdk/typescript/src/index";
import path from "path";

const binaryPath = path.resolve(process.cwd(), "bin/mime");

async function testAllFeatures() {
  console.log("🧪 MIME Feature Validation Test\n");
  console.log("Binary:", binaryPath);
  console.log("=".repeat(50));

  const browser = await MIME.connect({ binaryPath });
  const results: { tool: string; status: string; error?: string }[] = [];

  const testTool = async (name: string, fn: () => Promise<void>) => {
    process.stdout.write(`Testing ${name.padEnd(15)}... `);
    try {
      await fn();
      console.log("✅ PASS");
      results.push({ tool: name, status: "PASS" });
    } catch (e: any) {
      console.log(`❌ FAIL: ${e.message}`);
      results.push({ tool: name, status: "FAIL", error: e.message });
    }
  };

  // Use a simple, reliable test site
  const testUrl = "https://example.com";

  // 1. navigate
  await testTool("navigate", async () => {
    await browser.navigate(testUrl);
  });

  // 2. click
  await testTool("click", async () => {
    await browser.click("a");
  });

  // 3. type (using a pre-existing textarea on a form page)
  await testTool("type", async () => {
    // Use httpbin which has forms
    await browser.navigate("https://httpbin.org/forms/post");
    await browser.type('input[name="custname"]', "Test User");
  });

  // 4. extract
  await testTool("extract", async () => {
    await browser.navigate(testUrl);
    const text = await browser.extract("h1");
    if (!text) throw new Error("No text extracted");
  });

  // 5. screenshot (PNG)
  await testTool("screenshot", async () => {
    const data = await browser.screenshot();
    if (!data || !data.screenshot) throw new Error("No screenshot data");
  });

  // 6. execute (rod's Eval expects expression, not statement)
  await testTool("execute", async () => {
    // Rod's Eval needs just an expression, not a return statement
    const result = await browser.execute("document.title");
    if (!result) throw new Error("No result from execute");
  });

  // 7. html
  await testTool("html", async () => {
    const data = await browser.html();
    if (!data.html || data.html.length < 100) throw new Error("HTML too small");
  });

  // 8. wait_for
  await testTool("wait_for", async () => {
    await browser.navigate(testUrl);
    await browser.waitFor("h1");
  });

  // 9. scroll (correct signature: selector, x, y)
  await testTool("scroll", async () => {
    await browser.scroll(undefined, 0, 100);
  });

  // 10. hover
  await testTool("hover", async () => {
    await browser.hover("a");
  });

  // 11. markdown
  await testTool("markdown", async () => {
    const md = await browser.markdown();
    if (!md || !md.markdown || md.markdown.length < 10) throw new Error("Markdown too small");
  });

  // 12. metadata
  await testTool("metadata", async () => {
    const meta = await browser.metadata();
    if (!meta || !meta.title) throw new Error("No metadata");
  });

  // 13. links (returns { links: [...], count: n })
  await testTool("links", async () => {
    const result = await browser.links();
    if (!result || !Array.isArray(result.links)) throw new Error("Links not in expected format");
  });

  // 14. get_cookies (returns { cookies: [...], count: n })
  await testTool("get_cookies", async () => {
    const result = await browser.getCookies();
    if (!result || !Array.isArray(result.cookies)) throw new Error("Cookies not in expected format");
  });

  // 15. clear_cookies
  await testTool("clear_cookies", async () => {
    await browser.clearCookies();
  });

  // 16. observe
  await testTool("observe", async () => {
    await browser.navigate(testUrl);
    const obs = await browser.observe();
    if (!obs) throw new Error("No observation result");
  });

  // 17. act (simple action on example.com)
  await testTool("act", async () => {
    await browser.navigate(testUrl);
    await browser.act("click on More information link");
  });

  // 18. crawl (with required params)
  await testTool("crawl", async () => {
    const result = await browser.crawl(testUrl, {
      max_pages: 1,
      max_depth: 1,
      same_domain: true,
    });
    if (!result) throw new Error("No crawl result");
  });

  // 19. map
  await testTool("map", async () => {
    const sitemap = await browser.map(testUrl);
    if (!sitemap) throw new Error("No sitemap returned");
  });

  // Summary
  console.log("\n" + "=".repeat(50));
  console.log("📊 SUMMARY");
  console.log("=".repeat(50));

  const passed = results.filter((r) => r.status === "PASS").length;
  const failed = results.filter((r) => r.status === "FAIL").length;

  console.log(`Passed: ${passed}/${results.length}`);
  console.log(`Failed: ${failed}/${results.length}`);

  if (failed > 0) {
    console.log("\n❌ Failed Tests:");
    results
      .filter((r) => r.status === "FAIL")
      .forEach((r) => console.log(`  - ${r.tool}: ${r.error}`));
  }

  await browser.close();
  process.exit(failed > 0 ? 1 : 0);
}

testAllFeatures().catch((e) => {
  console.error("Fatal error:", e);
  process.exit(1);
});
