module Main where

import System.IO
import Data.List

main :: IO ()
main = do
    input <- readFile "day1-input.txt"
    let
        pairs = parseInput input
        (left, right) = unzip pairs
        sortedLeft = sort left
        sortedRight = sort right
        result = sum $ zipWith diff sortedLeft sortedRight

    putStrLn $ show result

diff x y = abs (x-y)
parseInput input = toListTupleInt $ toListTuple $ splitPair $ breakLine input
breakLine input = takeWhile (\l -> l /= "") $ lines input
splitPair lines = map (\l -> words l) lines
toListTuple lines = map (\[x,y] -> (x,y)) lines
toListTupleInt lines = map (\(x,y) -> (read x :: Int, read y :: Int)) lines

