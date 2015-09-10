module IndexSpec where

import Test.Hspec

import DropWiki


spec :: Spec
spec =

  describe "renderIndex" $
    it "renders the correct HTML" $
      let testWiki =
            Folder "root" [ Folder "level1" [ Folder "level2" [ File (Page "level2_file.md" "")
                                                              , File (Page "level2_other_file.md" "")
                                                              ]
                                            , File (Page "level1_file" "")
                                            ]
                          , File (Page "root_file.org" "")
                          ]
          renderedIndex =
            "<li><h1>root</h1></li>\n" ++
            "<li>root_file.org</li>\n" ++
            "<li><h2>level1</h2></li>\n" ++
            "<li>level1_file</li>\n" ++
            "<li><h3>level2</h3></li>\n" ++
            "<li>level2_file.md</li>\n" ++
            "<li>level2_other_file.md</li>\n"
      in renderIndex testWiki `shouldBe` renderedIndex
