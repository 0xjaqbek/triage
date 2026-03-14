# PWA Conversion Design

**Date:** 2026-02-25
**Status:** Approved

## Goal

Convert TRIAGE MCI web app into an installable Progressive Web App with full offline support on iOS, Android, and Windows. All existing functionality must be preserved.

## Decisions

- **Font strategy:** Remove Google Fonts, use system monospace stack
- **Icon source:** `triage.png` provided by user, generate PNG sizes from it
- **App name:** "TRIAGE" (short name on home screen)
- **Data persistence:** localStorage — auto-save/restore all state
- **Confirmation dialogs:** Required before any destructive action (reset, clear, new event)

## New Files

| File | Purpose |
|------|---------|
| `manifest.json` | PWA manifest — app name, icons, theme color, display mode |
| `sw.js` | Service worker — caches all assets for offline use |
| `icons/icon-192.png` | Android/general icon (192x192) |
| `icons/icon-512.png` | Android splash / Windows (512x512) |
| `icons/apple-touch-icon.png` | iOS home screen (180x180) |

## Changes to index.html

1. Remove Google Fonts link, replace with system monospace stack (`ui-monospace, 'SF Mono', 'Cascadia Code', 'Segoe UI Mono', Consolas, monospace`)
2. Add `<link rel="manifest">` + iOS meta tags (`apple-mobile-web-app-capable`, status bar style, `apple-touch-icon`)
3. Register service worker on page load
4. Add localStorage persistence — auto-save all state on every mutation, auto-restore on load
5. Add confirmation dialogs before destructive actions (new event, reset, clear data)
6. On load with existing data: prompt to resume or start fresh (with confirmation on "start fresh")

## Service Worker Strategy

- Cache-first for all local assets (HTML, manifest, icons)
- Precache everything on install — fully offline after first visit
- Version-based cache busting for updates

## localStorage Persistence

- Single key `triageAppState` storing JSON of all app state
- `saveState()` called after every mutation (add patient, dispatch, etc.)
- `loadState()` called on page load — if data exists, offer to resume
- Clear/reset requires user confirmation

## Platform Compatibility

- **Android:** Standard PWA install via Chrome "Add to Home Screen"
- **iOS:** `apple-mobile-web-app-capable` + `apple-touch-icon` + standalone display
- **Windows:** Installable via Edge/Chrome PWA install prompt

## What Stays the Same

- All triage wizard logic
- Dispatch system
- Report generation and clipboard copy
- Download function
- UI layout and dark theme styling
