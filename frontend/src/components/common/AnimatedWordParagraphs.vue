<script setup lang="ts">
import { computed } from 'vue'

import { useCurrentLocale } from '@/content/siteContent'

interface ParagraphTokens {
  id: string
  text: string
  tokens: Array<{
    text: string
    delay: number
  }>
}

const props = withDefaults(
  defineProps<{
    paragraphs: string[]
    revealed?: boolean
  }>(),
  {
    revealed: false
  }
)

const currentLocale = useCurrentLocale()

function buildTokens(text: string, locale: string): string[] {
  const tokens: string[] = []

  if (typeof Intl !== 'undefined' && 'Segmenter' in Intl) {
    const segmenter = new Intl.Segmenter(locale, { granularity: 'word' })
    for (const segment of segmenter.segment(text)) {
      const chunk = segment.segment
      if (!chunk) {
        continue
      }

      if (segment.isWordLike) {
        tokens.push(chunk)
        continue
      }

      if (tokens.length === 0) {
        tokens.push(chunk)
        continue
      }

      tokens[tokens.length - 1] += chunk
    }

    if (tokens.length > 0) {
      return tokens
    }
  }

  const fallbackSegments = text.match(/[\p{Letter}\p{Number}\p{Script=Han}]+(?:\s+)?|[^\s]/gu) ?? [text]

  for (const chunk of fallbackSegments) {
    if (tokens.length === 0) {
      tokens.push(chunk)
      continue
    }

    if (/^\s+$/.test(chunk) || /^[^\p{Letter}\p{Number}\p{Script=Han}]+$/u.test(chunk)) {
      tokens[tokens.length - 1] += chunk
      continue
    }

    tokens.push(chunk)
  }

  return tokens
}

const segmentedParagraphs = computed<ParagraphTokens[]>(() => {
  const locale = currentLocale.value
  let offset = 0

  return props.paragraphs.map((paragraph, paragraphIndex) => {
    const tokens = buildTokens(paragraph, locale).map((text, tokenIndex) => ({
      text,
      delay: offset + tokenIndex * 0.045
    }))

    offset += tokens.length * 0.045 + 0.16

    return {
      id: `${paragraphIndex}-${paragraph.slice(0, 18)}`,
      text: paragraph,
      tokens
    }
  })
})
</script>

<template>
  <div :class="['animated-word-paragraphs', { 'is-revealed': revealed }]">
    <p
      v-for="paragraph in segmentedParagraphs"
      :key="paragraph.id"
      class="animated-word-paragraphs__item"
      :aria-label="paragraph.text"
    >
      <span
        v-for="(token, tokenIndex) in paragraph.tokens"
        :key="`${paragraph.id}-${tokenIndex}`"
        class="animated-word-paragraphs__word"
        :style="{ '--word-delay': `${token.delay}s` }"
        aria-hidden="true"
      >
        {{ token.text }}
      </span>
    </p>
  </div>
</template>

<style scoped>
.animated-word-paragraphs {
  display: grid;
  gap: 1.1rem;
}

.animated-word-paragraphs__item {
  line-height: inherit;
}

.animated-word-paragraphs__word {
  display: inline-block;
  white-space: pre-wrap;
  opacity: 0;
  transform: translateY(0.95rem);
  filter: blur(6px);
  will-change: opacity, transform, filter;
}

.animated-word-paragraphs.is-revealed .animated-word-paragraphs__word {
  animation: wordLift 0.68s var(--transition-slow) both;
  animation-delay: var(--word-delay);
}

@keyframes wordLift {
  from {
    opacity: 0;
    transform: translateY(0.95rem);
    filter: blur(6px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
    filter: blur(0);
  }
}

@media (prefers-reduced-motion: reduce) {
  .animated-word-paragraphs__word {
    opacity: 1;
    transform: none;
    filter: none;
    will-change: auto;
  }
}
</style>
