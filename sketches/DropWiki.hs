module DropWiki where

import           Data.List (sort)

type Markdown = String
type HTML = String
type Path = String
type Title = String

data Page = Page Title Markdown deriving (Eq, Ord)
data WikiTree = File Page
              | Folder Title [WikiTree] deriving (Eq, Ord)

renderIndex :: WikiTree -> HTML
renderIndex = renderIndex' 1
              where wrapWithTag tag text = "<"++tag++">"++text++"</"++tag++">"
                    toLi = (++"\n") . wrapWithTag "li"
                    toHeader depth = wrapWithTag ("h" ++ show depth)

                    renderIndex' :: Int -> WikiTree -> HTML
                    renderIndex' _ (File (Page title _)) = toLi title
                    renderIndex' depth (Folder title children) =
                      renderedTitle ++ renderedChildren
                      where renderedTitle = toLi $ toHeader depth title
                            renderedChildren = concatMap (renderIndex' (depth + 1)) (sort children)

-- renderMD :: Markdown -> HTML
-- renderPage :: Page -> HTML

-- search :: WikiTree -> []SearchResult

-- serve :: Path -> IO ()
