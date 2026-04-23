import { computed } from 'vue'
import { createI18n } from 'vue-i18n'

import en from '@/i18n/messages/en'
import zhCN from '@/i18n/messages/zh-CN'

export const SUPPORTED_LOCALES = ['zh-CN', 'en'] as const
export type AppLocale = (typeof SUPPORTED_LOCALES)[number]

const STORAGE_KEY = 'mtn-locale'
const DEFAULT_LOCALE: AppLocale = 'zh-CN'

const messages = {
  'zh-CN': zhCN,
  en
} as const

function isSupportedLocale(value: string | null | undefined): value is AppLocale {
  return SUPPORTED_LOCALES.includes(value as AppLocale)
}

function resolveInitialLocale(): AppLocale {
  if (typeof window === 'undefined') {
    return DEFAULT_LOCALE
  }

  const stored = window.localStorage.getItem(STORAGE_KEY)
  if (isSupportedLocale(stored)) {
    return stored
  }

  const browserLocale = window.navigator.language
  if (browserLocale.toLowerCase().startsWith('zh')) {
    return 'zh-CN'
  }

  if (browserLocale.toLowerCase().startsWith('en')) {
    return 'en'
  }

  return DEFAULT_LOCALE
}

const initialLocale = resolveInitialLocale()

export const i18n = createI18n({
  legacy: false,
  locale: initialLocale,
  fallbackLocale: 'en',
  messages
})

function syncDocumentLanguage(locale: AppLocale): void {
  if (typeof document !== 'undefined') {
    document.documentElement.lang = locale
  }
}

syncDocumentLanguage(initialLocale)

export function setLocale(locale: AppLocale): void {
  if (!isSupportedLocale(locale)) {
    return
  }

  i18n.global.locale.value = locale
  syncDocumentLanguage(locale)

  if (typeof window !== 'undefined') {
    window.localStorage.setItem(STORAGE_KEY, locale)
  }
}

export function useCurrentLocale() {
  return computed(() => i18n.global.locale.value as AppLocale)
}
