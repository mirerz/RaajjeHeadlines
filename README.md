# 📡 Raajjé HEADLINES | Digital Ready v1.0

### **The Maldivian Agentic Newsroom Powered by 729 AI**

Raajjé HEADLINES is a high-intelligence, multi-platform news ecosystem designed for the modern Maldivian media landscape. It combines real-time **parallel scraping**, **AI-driven rephrasing**, and **hyper-vibrant 80s Cyber visuals** into a single production-ready powerhouse.

---

## 🚀 Key Ecosystem Features

### 🤖 1. The Agentic Aggregator
- **Concurrent Polling**: Automated 15-minute sync cycles across *Mihaaru*, *Sun*, *Vaguthu*, and *VNews*.
- **Autonomous Editor (AI)**: Rewrites headlines and body text to achieve 100% original Dhivehi content.
- **Subscriber Briefings**: Generates high-impact "Bullet Point Summaries" and "Deep-Dive Explainers" for every article.

### 🛠️ 2. Editorial Command Center
- **Cyan Cyber-Design**: High-visibility "Thought Monitoring" dashboard for staff review.
- **RBAC Security**: Multi-seat permissions for **Admins**, **Editors**, and **Subscribers**.
- **Social Gateway**: Integrated one-tap registration for **Gmail**, **Outlook**, **WhatsApp**, and **Telegram**.

### 📱 3. Mobile "Super" App (Flutter)
- **Vibrant Android Experience**: System-theme synced "Hyper-Color" 80s aesthetic.
- **Malé Utilities**: Integrated live **Prayer Times**, **Hijri Calendar**, and **Transit/Ferries** widgets.
- **Push Notifications**: Real-time breaking news alerts via Firebase (FCM).

---

## 🛠️ Infrastructure & Tech Stack
| Service | Technology | Role |
| :--- | :--- | :--- |
| **Backend API** | Go 1.23 (Fiber / GORM) | Central News Server & Auth |
| **Aggregator** | Go (Colly / ADK Go 1.0) | Automated News Ingestion |
| **Frontend** | Alpine.js / Tailwind CSS | Hyper-Vibrant WebApps |
| **Mobile App** | Flutter | Native Media Experience |
| **Database** | PostgreSQL 16 | News Records & User Data |
| **Cache** | Redis 7 | Live Feed & Session Sync |

---

## 🛰️ Deployment Guide (Docker)

To launch the full newsroom cluster:

1.  **Clone the Repo**:
    ```bash
    git clone [your-github-repo-url]
    cd raajje-headlines
    ```

2.  **Environment Setup**:
    Copy `.env.example` to `.env` and configure your AI API keys and database credentials.

3.  **One-Click Launch**:
    ```bash
    docker-compose up -d --build
    ```

4.  **Initial Access**:
    - **Dashboard**: `http://localhost:8080/dashboard`
    - **Live Feed**: `http://localhost:8080/`
    - **API**: `http://localhost:8080/api/v1`

---

### ✅ Mission Ready
**Project Version:** 1.0.0
**Project Lead:** 729 Holdings Agentic AI
**Status:** PRODUCTION READY

*Copyright © 2026 Raajjé HEADLINES. All Rights Reserved.*
