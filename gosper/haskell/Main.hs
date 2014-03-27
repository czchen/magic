import Data.List.Split

import Gosper

genList :: Integer -> Integer -> [Integer] -> [Integer]
genList item count list
    | count == 0 = list
    | otherwise  = item: genList (next item) (count - 1) list

main = do
    putStrLn "Enter [first number] [count]"
    line <- getLine
    let
        input = splitOn " " line
        item = read $ head input :: Integer
        count = read $ head $ tail input :: Integer
        in
            putStrLn $ show $ genList item count []
