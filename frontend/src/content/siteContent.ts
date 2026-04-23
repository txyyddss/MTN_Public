import { computed } from 'vue'

import { SUPPORTED_LOCALES, i18n, setLocale, useCurrentLocale } from '@/i18n'
import en from '@/i18n/messages/en'
import type { FixedLeaderboardType, StatGroup } from '@/types/api'
import type { CoreMember } from '@/types/coreMembers'

export type LocaleMessages = typeof en
export type SiteContent = LocaleMessages['siteContent']
export type NavItem = SiteContent['app']['nav'][number]
export type FeatureCardContent = SiteContent['home']['features'][number]
export type HeroContent = SiteContent['home']['hero']
type SiteFormatKey = keyof SiteContent['formats']

function getMessages(): LocaleMessages {
  return i18n.global.getLocaleMessage(i18n.global.locale.value) as LocaleMessages
}

function touchLocale(): void {
  void i18n.global.locale.value
}

function formatSitePattern(key: SiteFormatKey, values: Record<string, string | number> = {}): string {
  touchLocale()
  return i18n.global.t(`siteContent.formats.${key}`, values)
}

export function useSiteContent() {
  return computed(() => {
    touchLocale()
    return getMessages().siteContent
  })
}

export function useCoreMembers() {
  return computed(() => {
    touchLocale()
    return getMessages().coreMembers as unknown as CoreMember[]
  })
}

export function usePlayerStatGroups() {
  return computed(() => {
    touchLocale()
    return getMessages().playerStatGroups as unknown as StatGroup[]
  })
}

export function useLocaleOptions() {
  return computed(() => {
    touchLocale()
    return getMessages().siteContent.app.locale.options
  })
}

export function getLocaleValue(): string {
  return i18n.global.locale.value
}

export function getLeaderboardLabel(key: FixedLeaderboardType): string {
  return getMessages().leaderboardLabels[key] ?? key
}

export function getStatTranslation(key: string): string | undefined {
  return (getMessages().statTranslations as Record<string, string>)[key]
}

export function getStatCategoryLabel(key: string): string | undefined {
  return (getMessages().statCategories as Record<string, string>)[key]
}

export function getSkillLabel(key: string): string | undefined {
  return (getMessages().skillLabels as Record<string, string>)[key]
}

export function getPlatformLabel(platform: string): string {
  return getMessages().platformLabels[platform as keyof LocaleMessages['platformLabels']] ?? platform
}

export function getOreLabel(ore: string): string {
  const normalizedKey = ore.trim().toLowerCase().replace(/\s+/g, '_')
  return getMessages().oreLabels[normalizedKey as keyof LocaleMessages['oreLabels']] ?? ore
}

export function getAdvancementTypeLabel(type: 'task' | 'goal' | 'challenge'): string {
  return getMessages().advancementTypes[type] ?? type
}

export function getAdvancementCategoryLabel(key: string): string {
  return getMessages().advancementCategories[key as keyof LocaleMessages['advancementCategories']] ?? key.replace(/[_-]/g, ' ')
}

export function getAdvancementName(key: string, fallback: string): string {
  return (getMessages().advancementNames as Record<string, string>)[key] ?? fallback
}

export function formatPlayerCount(count: number): string {
  const localizedCount = count.toLocaleString(getLocaleValue())
  const nounKey: SiteFormatKey = count === 1 ? 'playerSingular' : 'playerPlural'
  return formatSitePattern('playerCount', { count: localizedCount, noun: formatSitePattern(nounKey) })
}

export function formatDurationHoursShort(value: string | number): string {
  return formatSitePattern('durationHoursShort', { value })
}

export function formatDurationHoursWide(value: string | number): string {
  return formatSitePattern('durationHoursWide', { value })
}

export function formatDurationMinutesShort(value: string | number): string {
  return formatSitePattern('durationMinutesShort', { value })
}

export function formatDistanceKilometers(value: string | number): string {
  return formatSitePattern('distanceKilometers', { value })
}

export function formatDistanceMeters(value: string | number): string {
  return formatSitePattern('distanceMeters', { value })
}

export function formatLeaderboardTitle(label: string): string {
  return formatSitePattern('leaderboardTitle', { label })
}

export { SUPPORTED_LOCALES, setLocale, useCurrentLocale }
