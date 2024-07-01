---
layout: default
title: AI Prompts
nav_order: 1
parent: Artificial Intelligence
grand_parent: Algorithms and Data Structures
permalink: /docs/algorithms-and-data-structures/ai/ai-prompts
---

- [go back to index]({% link docs/algorithms-data-structures/ai/index.md %})
- [go to prompt engineering]({% link docs/algorithms-data-structures/ai/prompt-engineering.md %})

## Rewrite and improve a text

Ditch “rewrite” and improve your AI content immediately with these prompts: [^1]

- __words better than “rewrite”__:
  - `paraphrase`: _this is useful when you want to avoid plagiarism._
  - `reframe`: _change the perspective or focus of the rewrite._
  - `summarize`: _when you want a quick overview of a lengthy topic._
  - `expand`: _for a more comprehensive understanding of a topic._
  - `explain`: _make the meaning of something clearer in the rewrite._
  - `reinterpret`: _provide a possible meaning or understanding._
  - `simplify`: _reduce the complexity of the language._
  - `elaborate`: _add more detail or explanation to a given point._
  - `amplify`: _strengthen the message or point in the rewrite._
  - `clarify`: _make a confusing point or statement clearer._
  - `adapt`: _modify the text for a different audience or purpose._
  - `modernize`: _update older language or concepts to be more current._
  - `formalize`: _this asks to rewrite informal or casual language into a more formal or professional style. Useful for business or academic contexts._
  - `informalize`: _use this for social media posts, blogs, email campaigns, or any context where a more colloquial style and relaxed tone is right._
  - `condense`: _make the rewrite shorter by restricting it to key points._
  - `emphasize/reiterate`: _highlight certain points more than others._
  - `diversify`: _add variety, perhaps in sentence structure or vocabulary._
  - `neutralize`: _remove bias or opinion, making the text more objective._
  - `streamline`: _remove unnecessary content or fluff._
  - `enrich/embellish`: _add more pizzazz or detail to the rewrite._
  - `illustrate`: _provide examples to better explain the point._
  - `synthesize`: _combine different pieces of information._
  - `sensationalize`: _make the rewrite more dramatic. Great for clickbait!_
  - `humanize`: _make the text more relatable or personal. Great for blogs!_
  - `elevate`: _prompt for a rewrite that is more sophisticated or impressive._
  - `illuminate`: _prompt for a rewrite that is crystal-clear or enlightening._
  - `enliven/energize`: _means make the text more lively or interesting._
  - `soft-pedal`: _means to downplay or reduce the intensity of the text._
  - `exaggerate`: _when you want to hype-up hyperbole in the rewrite. Great for sales pitches (just watch those pesky facts)!_
  - `downplay`: _when you want a more mellow, mild-mannered tone. Great for research, and no-nonsense evidence-based testimonials._
  - `glamorize`: _prompt to make the rewrite sexier and more appealing._
- [list of tones of voice](#list-of-tones-of-voice)
- __rewriting text using__
  - **_"reinterpret the above from an `X` perspective, for publication as `Y`"_**
    - `X` could be your intended audience (e.g. consumers, Americans, youth), or a certain mindset or ideology (business-oriented, environmentalist, etc.
    - `Y` is the type of content you’d like to create (e.g. eBook, article, email, blog, social media post, marketing material, etc). Modify and adapt the prompt!
  - **_"reuse the above concept in a `{insert}` tone of voice, that will resonate with `{insert intended audience}`"_**
    - another ‘rewrite’ that comes from [Amanda Weston at Blogs By Jasper](https://www.blogsbyjarvis.com/).
  - **_“rewrite the above in a tone of voice that is professional, informative, and friendly, emphasizing clarity, conciseness, objectivity, relevance, and organization in a conversational tone to engage the audience”_**

## Text Classification

- _"Classify the sentiment of the following text as either positive or negative."_
  - <details markdown="block"><summary><i>more...</i></summary>

    - | prompts  |
      | :------  |
      | "Categorize the following news article into one of these categories: Politics, Technology, Sports, Entertainment, or Business." |
      | "Assign relevant tags to the following text. Choose as many as appropriate from this list: Technology, AI, Business, Innovation, Social Media, Privacy, Cybersecurity." |
      | Classify the following text into one of these categories: <br> Technology, Business, or Politics. <br>Provide a confidence score (0-100) for each category. <br>""" <br> ... <br>""" <br>Classification: <br>Technology: [score] <br>Business: [score] <br>Politics: [score] |

    </details>

## Journal Writing Prompts

```
Output a journal writing prompt example - to help me write what I've done and
how I felt on such a day.  Give me insights into why I feel <my_mood> or any
other mood during the day.

##

<my_mood>: in good spirits
```
<sup><i>**list of moods**: positive, enthusiastic, happy, excited, cheerful, delighted,
thrilled, in high spirits, in good spirits, uplifting, heartening, fulfilling,
pleasing, delightful, pleasurable, enjoyable, agreeable, satisfying.</i></sup>

```
Output a journal writing prompt example - to help me write what I've done and
how I felt on such a day; one thing I can think of today is <one_today_event>.
Give me insights into why I felt <my_mood> or any other mood during the day.
Besides <one_today_event>, help me remember anything else.

##

<one_today_event>: a birthday party
<my_mood>: in good spirits
```


<details markdown="block"><summary>others things related to journal writing prompts</summary>

```
<one_today_event>: a birthday party
<my_mood>: positive, enthusiastic, happy, excited, cheerful, delighted,
thrilled, in high spirits, in good spirits, uplifting, heartening, fulfilling,
pleasing, delightful, pleasurable, enjoyable, agreeable, satisfying.
```

- <a id="list-of-tones-of-voice"></a> **tones of voice**
  - in action
    - happy
    - compelling
    - rapturous
  - writing assistant
    - ‘blog review’ _<sub>when to have to create a product review</sub>_
    - ‘copywriting’ _<sub>when to have a sales-like copy</sub>_
  - <details markdown="block"><summary>writing and marketing</summary>

      - simple
      - interesting
      - informative
      - creative
      - compelling
      - copywriting
      - Neil Patel
      - Guy Kawasaki
      - Gary Vaynerchuk
    </details>
  - <details markdown="block"><summary>serious, journalistic, academic</summary>

      - New York Times
      - The Guardian
      - Scientific American
      - factual
      - informative
      - descriptive
      - research
      - Harvard Business Review
      - academic
      - serious
      - somber
      - journalistic
      - blog review
      - court
      - legal
    </details>
  - <details markdown="block"><summary>happy and joyful</summary>

      - ecstatic
      - positive
      - motivational
      - joyful
      - playful
      - passionate
      - romantic
      - affirmative
      - enthusiastic
      - uplifting
      - oprah
      - creative
    </details>
  - <details markdown="block"><summary>positive emotional</summary>

      - motivated
      - inspired
      - lovestruck
      - enchanted
      - blissful
      - excited
      - thrilling
      - amazing/amazed
      - powerful
      - light
    </details>
  - <details markdown="block"><summary>negative emotional</summary>

      - anxious
      - melancholy
      - depressed
      - sad
      - infuriating
      - disturbing
      - angry
      - hopeless
      - enraged
      - provocative
      - embarrassing
      - greedy
      - dark
      - sad
      - abusive
    </details>
  - <details markdown="block"><summary>other tone of voices to consider</summary>

      - approachable
      - excited
      - playful
      - assertive
      - formal
      - poetic
      - bold
      - friendly
      - positive
      - candid
      - funny
      - powerful
      - caring
      - gentle
      - professional
      - casual
      - helpful
      - quirky
      - cheerful
      - hopeful
      - reassuring
      - clear
      - humorous
      - reflective
      - commanding
      - informal
      - respectful
      - comprehensive
      - informative
      - romantic
      - confident
      - inspirational
      - sarcastic
      - conversational
      - inspiring
      - scientific
      - curious
      - lively
      - serious
      - detailed
      - melancholic
      - technical
      - educational
      - motivational
      - thought-provoking
      - eloquent
      - negative
      - thoughtful
      - emotional
      - neutral
      - uplifting
      - empathetic
      - nostalgic
      - urgent
      - empowering
      - offbeat
      - vibrant
      - encouraging
      - passionate
      - visionary
      - engaging
      - personal
      - witty
      - enthusiastic
      - persuasive
      - zealous
    </details>
  - <details markdown="block"><summary>diverse personalities</summary>

      - knowledgeable companion
      - approachable educator
      - visual narrator
      - fairytale friendliness
      - witty cynic
      - empathetic encourager
      - playful entertainer
      - confident commander
      - thoughtful advisor
      - provocative prankster.
    </details>
  - miscellaneous
    - constructive / feedback :: _provide constructive feedback_
      - make sure that you explain the "Why" behind the feedback so that it reduces the need for follow-ups and gives the necessary context.

-----
<!-- others things related to journal writing prompts -->
</details>

----

[^1]: [31 AI Prompts better than “Rewrite”](https://medium.com/the-generator/31-ai-prompts-better-than-rewrite-b3268dfe1fa9)
