import { describe, it, expect, mock, beforeEach } from "bun:test";
import { MIME } from "./mime";
import { MCPClient } from "./client";

// Mock callTool function
const mockCallTool = mock(() => Promise.resolve({
    content: [{ type: 'text', text: '{}' }]
}));

// Mock MCPClient
const mockClient = {
    connect: mock(() => Promise.resolve()),
    callTool: mockCallTool,
    disconnect: mock(() => Promise.resolve()),
    isConnected: () => true
} as unknown as MCPClient;

describe("MIME SDK", () => {
    let browser: MIME;

    beforeEach(() => {
        // Reset mocks
        mockCallTool.mockClear();
        // Create browser instance with mocked client (bypassing private constructor)
        browser = new (MIME as any)(mockClient);
    });

    it("should navigate", async () => {
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify({
                url: "https://example.com",
                title: "Example",
                httpStatus: 200
            })}]
        });

        const result = await browser.navigate("https://example.com");

        expect(mockCallTool).toHaveBeenCalledWith("navigate", { url: "https://example.com" });
        expect(result.url).toBe("https://example.com");
    });

    it("should click", async () => {
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify({ success: true }) }]
        });

        await browser.click("#btn");

        expect(mockCallTool).toHaveBeenCalledWith("click", { selector: "#btn" });
    });

    it("should type", async () => {
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify({ success: true }) }]
        });

        await browser.type("#input", "hello");

        expect(mockCallTool).toHaveBeenCalledWith("type", { selector: "#input", text: "hello" });
    });

    it("should extract", async () => {
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify({ text: "Extracted Text" }) }]
        });

        const text = await browser.extract("h1");

        expect(mockCallTool).toHaveBeenCalledWith("extract", { selector: "h1" });
        expect(text).toBe("Extracted Text");
    });

    it("should observe", async () => {
        const mockObs = {
            url: "https://example.com",
            forms: [],
            clickable: [],
            inputs: []
        };
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify(mockObs) }]
        });

        const obs = await browser.observe();

        expect(mockCallTool).toHaveBeenCalledWith("observe", {});
        expect(obs.url).toBe("https://example.com");
    });

    it("should act", async () => {
        mockCallTool.mockResolvedValueOnce({
            content: [{ type: 'text', text: JSON.stringify({ success: true, action: "click" }) }]
        });

        const res = await browser.act("click login");

        expect(mockCallTool).toHaveBeenCalledWith("act", { instruction: "click login" });
        expect(res.success).toBe(true);
    });
});
