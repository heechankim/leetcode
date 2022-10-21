class Solution:
    def maxProfit(self, prices: List[int]) -> int:

        price_min = 10000
        profit = 0

        for price in prices:

            price_min = min(price, price_min)

            profit = max(profit, price - price_min)

        return profit