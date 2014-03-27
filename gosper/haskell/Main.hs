import Data.List.Split
import System.Environment

import Gosper

genList :: Integer -> Integer -> [Integer] -> [Integer]
genList item count list
    | count == 0 = list
    | otherwise  = item: genList (next item) (count - 1) list

main = do
    args <- getArgs
    let
        item = read $ head args :: Integer
        count = read $ head $ tail args :: Integer
        in
            putStrLn $ show $ genList item count []
