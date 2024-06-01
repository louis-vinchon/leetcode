package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/replace-words/description/

// In English, we have a concept called root, which can be followed by some
// other word to form another longer word - let's call this word derivative.
// For example, when the root "help" is followed by the word "ful", we can form
// a derivative "helpful".
//
// Given a dictionary consisting of many roots and a sentence consisting of
// words separated by spaces, replace all the derivatives in the sentence with
// the root forming it. If a derivative can be replaced by more than one root,
// replace it with the root that has the shortest length.
//
// Return the sentence after the replacement.

//Brute force approach.
func replaceWordsBruteForce(dictionary []string, sentence string) string {
  // Remove duplicates or otherwise redundant roots.
  cleanRoots := []string{}
  ignore := make(map[int]bool)
  for i := 0 ; i < len(dictionary) ; i++ {
    if ignore[i] {
      // fmt.Printf("ignore\n")
      continue
    }

    shortest := i

    for k := i + 1 ; k < len(dictionary) ; k++ {
      // fmt.Printf("i %d %s | k %d %s\n", i, dictionary[i], k, dictionary[k])
      if ignore[k] {
        // fmt.Printf("ignore\n")
        continue
      }

      localShortest := shortest
      longest := k
      // fmt.Printf("before sort: %s %s\n", dictionary[localShortest], dictionary[longest])
      if len(dictionary[localShortest]) > len(dictionary[longest]) {
        localShortest, longest = longest, localShortest
      }
      // fmt.Printf("after  sort: %s %s\n", dictionary[localShortest], dictionary[longest])

      if dictionary[longest][:len(dictionary[localShortest])] == dictionary[localShortest] {
        // fmt.Printf("short %s\nlong %s\n---\n", dictionary[localShortest], dictionary[longest])
        shortest = localShortest
        ignore[longest] = true
        ignore[localShortest] = true
      }
    }
    cleanRoots = append(cleanRoots, dictionary[shortest])
  }

  // fmt.Printf("ignore %v\n", ignore)
  // fmt.Printf("roots %v\n", dictionary)
  // fmt.Printf("clean roots %v\n", cleanRoots)

  words := strings.Split(sentence, " ")
  rootsPerWord := make(map[string]string)

  for _, word := range words {
    // fmt.Printf("---\nword %s\n", word)
    _, alreadyProcessed := rootsPerWord[word]
    if alreadyProcessed {
      // fmt.Printf("processed already\n")
      continue
    }

    // Match against each root.
    for _, root := range cleanRoots {
      // fmt.Printf("root %s\n", root)
      if len(word) < len(root) {
        // fmt.Printf("word too short\n")
        continue
      }
      if word[:len(root)] != root {
        // fmt.Printf("root doesn't match\n")
        continue
      }

      // No need to check for shorter roots as we have already eliminated the longer ones.
      rootsPerWord[word] = root
    }
  }
  // fmt.Printf("roots %v\n", rootsPerWord)



  newSentence := ""
  for i := 0 ; i < len(words) ; i++ {
    root, ok := rootsPerWord[words[i]]
    if ok {
      newSentence += " " + root
    } else {
      newSentence += " " + words[i]
    }
  }

  // Don't forget to trim the leading space.
  return newSentence[1:]
}

// ------------------------------------------

type HashTrie struct {
  children map[rune]*HashTrie
  isWord bool
}

func newHashTrie() *HashTrie {
  return &HashTrie{
    children: make(map[rune]*HashTrie),
    isWord: false,
  }
}

func (trie *HashTrie) Add(word string) {
  if len(word) == 0 {
    trie.isWord = true
    return
  }

  _, ok := trie.children[rune(word[0])]
  if !ok {
    trie.children[rune(word[0])] = newHashTrie()
  }
  childTrie := trie.children[rune(word[0])]
  childTrie.Add(word[1:])
}

func (trie HashTrie) FindWord(word string) bool {
  if len(word) == 0 {
    return trie.isWord
  }

  childTrie, ok := trie.children[rune(word[0])]
  if !ok {
    return false
  }

  return childTrie.FindWord(word[1:])
}

func (trie HashTrie) FindPrefix(word string) bool {
  if len(word) == 0 {
    return true
  }

  childTrie, ok := trie.children[rune(word[0])]
  if !ok {
    return false
  }

  return childTrie.FindPrefix(word[1:])
}

func (trie HashTrie) GetShortestRoot(word string) int {
  if len(word) == 0 {
    return 0
  }
  if trie.isWord {
    return 0
  }

  childTrie, ok := trie.children[rune(word[0])]
  if !ok {
    return len(word)
  }

  return 1 + childTrie.GetShortestRoot(word[1:])
}

// func (trie Trie) String() string {
//   s := ""
//   for char, childTrie := range trie.children {
//     childString := childTrie.String()
//     s = string(char) + childString
//   }
//
//   return s
// }

// ---------------------

type ArrTrie struct {
  // Only lower case letters.
  children [26]*ArrTrie
  isWord bool
}

func newArrTrie() *ArrTrie {
  return &ArrTrie{
    children: [26]*ArrTrie{},
    isWord: false,
  }
}

func (trie *ArrTrie) Add(word string) {
  if len(word) == 0 {
    trie.isWord = true
    return
  }

  childTrie := trie.children[word[0] - 'a']
  if childTrie == nil {
    trie.children[word[0] - 'a'] = newArrTrie()
  }
  childTrie = trie.children[word[0] - 'a']
  childTrie.Add(word[1:])
}

func (trie ArrTrie) FindWord(word string) bool {
  if len(word) == 0 {
    return trie.isWord
  }

  childTrie := trie.children[word[0] - 'a']
  if childTrie == nil {
    return false
  }

  return childTrie.FindWord(word[1:])
}

func (trie ArrTrie) FindPrefix(word string) bool {
  if len(word) == 0 {
    return true
  }

  childTrie := trie.children[word[0] - 'a']
  if childTrie == nil {
    return false
  }

  return childTrie.FindPrefix(word[1:])
}

func (trie ArrTrie) GetShortestRoot(word string) int {
  if len(word) == 0 {
    return 0
  }
  if trie.isWord {
    return 0
  }

  childTrie := trie.children[word[0] - 'a']
  if childTrie == nil {
    return len(word)
  }

  return 1 + childTrie.GetShortestRoot(word[1:])
}



func replaceWordsHashTrie(dictionary []string, sentence string) string {
  trie := newHashTrie()
  for _, root := range dictionary {
    trie.Add(root)
  }

  newString := ""
  words := strings.Split(sentence, " ")
  for _, word := range words {
    length := trie.GetShortestRoot(word)

    newString += " " + word[:length]
  }

  return newString[1:]
}

func replaceWordsArrTrie(dictionary []string, sentence string) string {
  trie := newArrTrie()
  for _, root := range dictionary {
    trie.Add(root)
  }

  newString := ""
  words := strings.Split(sentence, " ")
  for _, word := range words {
    length := trie.GetShortestRoot(word)

    newString += " " + word[:length]
  }

  return newString[1:]
}

func main() {
  // replaceWords := replaceWordsBruteForce
  replaceWords := replaceWordsArrTrie

  dictionary := []string{"cat","bat","rat"}
  sentence := "the cattle was rattled by the battery"
  fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))

  // dictionary = []string{"a","b","c"}
  // sentence = "aadsfasf absbs bbab cadsfafs"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))
  //
  // dictionary = []string{"rat","ra"}
  // sentence = "ratify"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))
  //
  // dictionary = []string{"rat","ra"}
  // sentence = "ratify ratify"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))
  //
  // dictionary = []string{"ratif","rati","rat","ra"}
  // sentence = "ratify ratify"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))
  //
  // dictionary = []string{"ratif","rati","rat","ra","ra","ra","ra","ra","ra","ra","ra","ra"}
  // sentence = "ratify ratify"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))

  // dictionary = []string{"e","k","c","harqp","h","gsafc","vn","lqp","soy","mr","x","iitgm","sb","oo","spj","gwmly","iu","z","f","ha","vds","v","vpx","fir","t","xo","apifm","tlznm","kkv","nxyud","j","qp","omn","zoxp","mutu","i","nxth","dwuer","sadl","pv","w","mding","mubem","xsmwc","vl","farov","twfmq","ljhmr","q","bbzs","kd","kwc","a","buq","sm","yi","nypa","xwz","si","amqx","iy","eb","qvgt","twy","rf","dc","utt","mxjfu","hm","trz","lzh","lref","qbx","fmemr","gil","go","qggh","uud","trnhf","gels","dfdq","qzkx","qxw"}
  // sentence = "ikkbp miszkays wqjferqoxjwvbieyk gvcfldkiavww vhokchxz dvypwyb bxahfzcfanteibiltins ueebf lqhflvwxksi dco kddxmckhvqifbuzkhstp wc ytzzlm gximjuhzfdjuamhsu gdkbmhpnvy ifvifheoxqlbosfww mengfdydekwttkhbzenk wjhmmyltmeufqvcpcxg hthcuovils ldipovluo aiprogn nusquzpmnogtjkklfhta klxvvlvyh nxzgnrveghc mpppfhzjkbucv cqcft uwmahhqradjtf iaaasabqqzmbcig zcpvpyypsmodtoiif qjuiqtfhzcpnmtk yzfragcextvx ivnvgkaqs iplazv jurtsyh gzixfeugj rnukjgtjpim hscyhgoru aledyrmzwhsz xbahcwfwm hzd ygelddphxnbh rvjxtlqfnlmwdoezh zawfkko iwhkcddxgpqtdrjrcv bbfj mhs nenrqfkbf spfpazr wrkjiwyf cw dtd cqibzmuuhukwylrnld dtaxhddidfwqs bgnnoxgyynol hg dijhrrpnwjlju muzzrrsypzgwvblf zbugltrnyzbg hktdviastoireyiqf qvufxgcixvhrjqtna ipfzhuvgo daee r nlipyfszvxlwqw yoq dewpgtcrzausqwhh qzsaobsghgm ichlpsjlsrwzhbyfhm ksenb bqprarpgnyemzwifqzz oai pnqottd nygesjtlpala qmxixtooxtbrzyorn gyvukjpc s mxhlkdaycskj uvwmerplaibeknltuvd ocnn frotscysdyclrc ckcttaceuuxzcghw pxbd oklwhcppuziixpvihihp"
  // fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))

  dictionary = []string{"harqp","h","ha"}
  sentence = "harq"
  fmt.Printf("%v\n\n", replaceWords(dictionary, sentence))

  // trie := newTrie()
  // trie.Add("ab")
  // trie.Add("abc")
  // trie.Add("c")
  //
  // fmt.Println(trie)
  // fmt.Printf("prefix %s %t\n", "tool", trie.FindPrefix("tool"))
  // fmt.Printf("word %s %t\n", "tool", trie.FindWord("tool"))
  //
  // fmt.Printf("prefix %s %t\n", "abra", trie.FindPrefix("abra"))
  // fmt.Printf("word %s %t\n", "abra", trie.FindWord("abra"))
  //
  // fmt.Printf("prefix %s %t\n", "abracadabra", trie.FindPrefix("abracadabra"))
  // fmt.Printf("word %s %t\n", "abracadabra", trie.FindWord("abracadabra"))
  //
  // fmt.Printf("shortest %d\n", trie.GetShortestRoot("dba"))
}

