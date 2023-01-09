import isZero from "../random/sandbox";

describe("isZero", () => {
    it("returns true for 0", () => {
        expect(isZero(0)).toBe(true);
    });
    it("returns false for 1", () => {
        expect(isZero(1)).toBe(false);
    });
});