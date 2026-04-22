# Test Plan: Initial Epic Template

> **Epic**: [EPIC-1000 — Epic Title](../epics/EPIC-1000/EPIC-1000.md)
> **PRD**: [PRD](../epics/EPIC-1000/PRD.md)
> **Tech Design**: [Tech Design](../epics/EPIC-1000/TECH-DESIGN.md)
> Copy to `docs/sdlc/epics/EPIC-1000/TEST-PLAN.md`

---

## Metadata

| Field | Value |
|-------|-------|
| **Epic Key** | EPIC-1000 |
| **Author** | |
| **QA Reviewer** | |
| **Status** | `draft` / `review` / `approved` |
| **Created** | 2026-04-22 |

---

## 1. Test Scope

### In Scope
_What this test plan covers (derived from PRD acceptance criteria)._

| AC ID | Acceptance Criteria | Test Type |
|-------|-------------------|-----------|
| EPIC-1000-AC01 | (from PRD) | Unit / UI / Integration |
| EPIC-1000-AC02 | | |

### Out of Scope
_What is NOT tested and why._

---

## 2. Device & OS Matrix

> Mark required devices. Camera/hardware features MUST test on real device.

| Device | iOS | Real Device | Simulator | Priority |
|--------|-----|-------------|-----------|----------|
| iPhone SE 3rd | 16.6 (minimum) | ⬜ | ⬜ | Must |
| iPhone 14 | 17.x | ⬜ | ⬜ | Must |
| iPhone 15 Pro | 18.x | ⬜ | ⬜ | Must |
| iPhone 16 Pro Max | 18.x | ⬜ | ⬜ | Should |

### Simulator Limitations (cannot test)
- Camera capture, focus, flash, exposure
- Push notifications
- NFC, Face ID hardware
- Real network transitions (WiFi ↔ cellular)
- Background upload with real URLSession background config

---

## 3. Unit Tests

| ID | Component | Test Description | File | Status |
|----|-----------|-----------------|------|--------|
| EPIC-1000-UT01 | ViewModel | State transitions: idle → loading → success/error | `DreemCatcherTests/` | ⬜ |
| EPIC-1000-UT02 | ViewModel | Data transformation / mapping | | ⬜ |
| EPIC-1000-UT03 | Service | API call with mocked HTTPClient | | ⬜ |
| EPIC-1000-UT04 | Model | Codable encode/decode with full response | | ⬜ |
| EPIC-1000-UT05 | Model | Codable decode with missing optional fields | | ⬜ |
| EPIC-1000-UT06 | Model | Codable decode with unknown extra fields | | ⬜ |

---

## 4. UI Tests

| ID | Flow | Steps | Expected | Device | Status |
|----|------|-------|----------|--------|--------|
| EPIC-1000-UI01 | Happy path | 1. Open feature 2. Do action 3. See result | Result displayed | Simulator | ⬜ |
| EPIC-1000-UI02 | Error state | 1. Trigger error condition | Error UI shown | Simulator | ⬜ |
| EPIC-1000-UI03 | Empty state | 1. Open with no data | Empty state shown | Simulator | ⬜ |

---

## 5. Integration Tests

| ID | Flow | Components | Precondition | Expected | Status |
|----|------|-----------|--------------|----------|--------|
| EPIC-1000-IT01 | End-to-end | View → VM → Service → API | Authenticated, DEV backend | Data loaded | ⬜ |
| EPIC-1000-IT02 | Token refresh | Expired token → interceptor → retry | Expired token | Transparent refresh | ⬜ |

---

## 6. Mobile-Specific Tests

### Camera Tests (Real Device Only)

| ID | Scenario | Device | Status |
|----|----------|--------|--------|
| EPIC-1000-CAM01 | Camera opens, preview shows | Real device | ⬜ |
| EPIC-1000-CAM02 | Tap to focus at various points | Real device | ⬜ |
| EPIC-1000-CAM03 | Volume button capture | Real device | ⬜ |
| EPIC-1000-CAM04 | Camera permission denied → settings prompt | Simulator OK | ⬜ |

### Network Tests

| ID | Scenario | How to Test | Expected | Status |
|----|----------|-------------|----------|--------|
| EPIC-1000-NET01 | Offline: no network | Airplane mode | Graceful error / cached data | ⬜ |
| EPIC-1000-NET02 | Network loss mid-upload | Toggle airplane during upload | Retry / resume | ⬜ |
| EPIC-1000-NET03 | WiFi → cellular transition | Switch network | Socket reconnects | ⬜ |
| EPIC-1000-NET04 | Slow network (2G) | Network Link Conditioner | Timeout handling | ⬜ |

### App Lifecycle Tests

| ID | Scenario | Steps | Expected | Status |
|----|----------|-------|----------|--------|
| EPIC-1000-LC01 | Background during feature | Press Home | State preserved | ⬜ |
| EPIC-1000-LC02 | Return from background | Open app again | Resume correctly | ⬜ |
| EPIC-1000-LC03 | Memory warning | Simulate in Xcode | No crash, graceful dealloc | ⬜ |
| EPIC-1000-LC04 | Incoming call during capture | Call during camera | Session paused, resumes | ⬜ |
| EPIC-1000-LC05 | App killed and relaunched | Force kill | Clean restart, no stale state | ⬜ |

### Permission Tests

| ID | Permission | Scenario | Expected | Status |
|----|-----------|----------|----------|--------|
| EPIC-1000-PM01 | Camera | First time: allow | Camera works | ⬜ |
| EPIC-1000-PM02 | Camera | First time: deny | Show settings prompt | ⬜ |
| EPIC-1000-PM03 | Photo Library | Allow limited access | Show limited photos | ⬜ |

---

## 7. Performance Tests

| ID | Metric | Threshold | How to Measure | Status |
|----|--------|-----------|----------------|--------|
| EPIC-1000-PF01 | Screen load time | < 300ms | Instruments: Time Profiler | ⬜ |
| EPIC-1000-PF02 | Memory footprint | < 20MB increase | Instruments: Allocations | ⬜ |
| EPIC-1000-PF03 | No memory leaks | 0 leaks after 10 cycles | Instruments: Leaks | ⬜ |

---

## 8. Accessibility Tests

| ID | Check | Expected | Status |
|----|-------|----------|--------|
| EPIC-1000-A11Y01 | VoiceOver navigation | All elements announced | ⬜ |
| EPIC-1000-A11Y02 | Dynamic Type (largest) | Text scales, no truncation | ⬜ |
| EPIC-1000-A11Y03 | Color contrast | Meets WCAG AA (4.5:1) | ⬜ |

---

## 9. Regression Checklist

> Existing features that MUST still work after this change.

| Area | Quick Smoke Test | Status |
|------|-----------------|--------|
| Login / Auth | Can log in and reach dashboard | ⬜ |
| Camera capture | Can take photo and preview | ⬜ |
| Upload | Can upload image successfully | ⬜ |
| Chat / AI | Can send message and receive AI response | ⬜ |
| Session list | Sessions load with pagination | ⬜ |

---

## 10. Test Results Summary

> Fill after test execution.

| Category | Total | Pass | Fail | Blocked | Skip |
|----------|-------|------|------|---------|------|
| Unit | | | | | |
| UI | | | | | |
| Integration | | | | | |
| Mobile-specific | | | | | |
| Performance | | | | | |
| Accessibility | | | | | |
| Regression | | | | | |
| **Total** | | | | | |

**Overall Verdict**: ⬜ Pass / ⬜ Fail / ⬜ Pass with known issues

**Known Issues**:
| ID | Description | Severity | Ticket |
|----|-------------|----------|--------|
| | | | EPIC-1000 |

---

## 11. Sign-off

| Role | Name | Date | Approved |
|------|------|------|----------|
| QA Lead | | | ⬜ |
| Tech Lead | | | ⬜ |
| PM | | | ⬜ |
