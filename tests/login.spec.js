const { test, expect } = require("@playwright/test");

test("Log in using Universal Login", async ({ page }) => {
  await page.goto("/");
  await page.locator("#qsLoginBtn").waitFor({ timeout: 10_000 });
  await page.locator("#qsLoginBtn").click();

  await page.locator('input[name="username"]').waitFor({ timeout: 10_000 });
  await page.locator('input[name="username"]').fill("asdasd");
  await page.locator('input[name="password"]').fill("asdasd");
  await page.locator('button[type="submit"]').click();

  await expect(page.locator("#qsLogoutBtn")).toBeVisible({ timeout: 15_000 });
});
