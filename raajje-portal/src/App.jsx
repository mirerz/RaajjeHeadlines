import { useEffect, useState } from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import { 
  onAuthStateChanged, signInWithPopup, GoogleAuthProvider, signOut,
  RecaptchaVerifier, signInWithPhoneNumber
} from 'firebase/auth';
import { doc, getDoc, setDoc } from 'firebase/firestore';
import { auth, db } from './firebase';
import './index.css';
import { getHotspotsForEndheriOdi } from './dataconnect';
import ThinkingAnimation from './ThinkingAnimation';

function App() {
  const [user, setUser] = useState(null);
  const [role, setRole] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, async (u) => {
      setUser(u);
      if (u) {
        const userDoc = await getDoc(doc(db, "users", u.uid));
        if (userDoc.exists()) {
          setRole(userDoc.data().role);
        }
      }
      setLoading(false);
    });
    return () => unsubscribe();
  }, []);

  if (loading) return <div className="loading-nexus">HANDSHAKING...</div>;

  return (
    <Routes>
      <Route path="/" element={<PortalHome user={user} role={role} />} />
      <Route path="/editorial" element={<EditorialNexus user={user} role={role} />} />
      <Route path="/editorial/:sector" element={<EditorialNexus user={user} role={role} />} />
    </Routes>
  );
}

function PortalHome({ user, role }) {
  const [hotspots, setHotspots] = useState([]);
  const [feed, setFeed] = useState([]);
  const [activeCategory, setActiveCategory] = useState("Fahuge");

  const categories = [
    { id: "Fahuge", label: "Fahuge / (Landing Page)", icon: "🏠" },
    { id: "Politics", label: "Fahuge /Politics&Gorvernance", icon: "🏛️" },
    { id: "Environment", label: "Fahuge /Environment&Devolopment", icon: "🌴" },
    { id: "Business", label: "Fahuge /Bussiness&Economy", icon: "📊" },
    { id: "Youth", label: "Fahuge /Youth&Sports", icon: "⚡" },
    { id: "Women", label: "Fahuge /Women&Kids", icon: "🫂" },
    { id: "Geopolitics", label: "ފަހުގެ / ޖިއޯޕޮލިޓިކްސް", icon: "🌍" },
    { id: "Middle-East", label: "ފަހުގެ / މެދުއިރުމަތި", icon: "🏙️" },
    { id: "Asia", label: "ފަހުގެ / އޭޝިއާ", icon: "🌏" },
  ];

  useEffect(() => {
    async function loadHotspots() {
      try {
        const response = await getHotspotsForEndheriOdi();
        if (response.data?.loreHotspots) {
          setHotspots(response.data.loreHotspots);
        }
      } catch (e) {
        console.error("Oivaru Engine sync error:", e);
      }
    }
    async function loadFeed() {
      // Logic: Use the Vite environment variable for production, or fallback to local sandbox
      const baseUrl = import.meta.env.VITE_API_URL || "http://127.0.0.1:3005";
      try {
        const res = await fetch(`${baseUrl}/api/feed`);
        if (res.ok) {
          const data = await res.json();
          setFeed(data);
        }
      } catch (e) {
        console.error("Feed sync error:", e);
      }
    }
    loadHotspots();
    loadFeed();
  }, []);

  const heroArticle = feed.length > 0 ? feed[0] : null;
  const regularFeed = feed.slice(1);

  return (
    <>
      <header className="ambient-header">
        <div className="logo">
          adhu.space<span className="dot">.</span>
        </div>
        <div className="vibe-indicator">
          {user ? (
            <Link to="/editorial" className="editorial-link">
              {role === 'SUBSCRIBER' ? 'THE VAULT 🏺' : 'EDITORIAL HUB 🏮'}
            </Link>
          ) : (
            <Link to="/editorial" className="login-btn">JOIN THE REPUBLIC</Link>
          )}
        </div>
      </header>

      <nav className="category-bar">
        <div className="category-scroll">
          {categories.map((cat) => (
            <button
              key={cat.id}
              className={`category-tab ${activeCategory === cat.id ? 'active' : ''}`}
              onClick={() => setActiveCategory(cat.id)}
            >
              <span className="tab-icon">{cat.icon}</span>
              <span className="tab-label">{cat.label}</span>
            </button>
          ))}
        </div>
      </nav>

      {heroArticle ? (
        <section className="hero-section" style={{ backgroundImage: heroArticle.VisualURL ? `linear-gradient(to top, rgba(15, 15, 26, 1), rgba(15, 15, 26, 0.4)), url(${heroArticle.VisualURL})` : undefined, backgroundSize: 'cover', backgroundPosition: 'center' }}>
          <div className="hero-badges" style={{ display: 'flex', gap: '8px', marginBottom: '16px' }}>
            {heroArticle.IsBreaking && <span className="bg-red-600 text-white text-xs font-bold px-2 py-1 rounded animate-pulse">⚡ BREAKING</span>}
            {heroArticle.IsVerifiedGov && <span className="bg-cyan-600 text-white text-xs font-bold px-2 py-1 rounded">🛡️ VERIFIED GOV</span>}
            <span className="bg-purple-700 text-white text-xs font-bold px-2 py-1 rounded">{heroArticle.Category}</span>
          </div>
          <h1 className="hero-title rtl" style={{ fontSize: '2.5rem', lineHeight: '1.4' }}>
            {heroArticle.RephrasedHeadlineDv}
          </h1>
          
          <div className="mt-4 bg-[#1A1A2E]/80 p-4 rounded backdrop-blur-sm border border-[#4B0082]">
            <h3 className="text-cyan-400 text-xs uppercase tracking-widest mb-2 font-bold">INTELLIGENCE BRIEFING</h3>
            <ul className="rtl space-y-2">
              {(JSON.parse(heroArticle.SummaryBulletsDv || '[]')).map((bullet, i) => (
                 <li key={i} className="text-sm flex gap-2"><span className="text-cyan-400">✦</span> <span>{bullet}</span></li>
              ))}
            </ul>
          </div>

          <div className="oivaru-tag mt-6 inline-block">
            <span>✨</span> Synthesized from {heroArticle.SourceName}
          </div>
        </section>
      ) : (
        <section className="hero-section flex items-center justify-center">
          <ThinkingAnimation message="Synthesizing Hero Article..." />
        </section>
      )}

      <section className="subscriber-threshold-section">
        <SubscriberThreshold />
      </section>

      <section className="feed-grid p-6" style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(320px, 1fr))', gap: '24px' }}>
        {regularFeed.map(article => (
          <div key={article.ID} className="bg-[#1A1A2E] border border-[#4B0082] rounded-xl overflow-hidden shadow-lg hover:shadow-cyan-500/20 transition transform hover:-translate-y-1">
            {article.VisualURL && (
              <img src={article.VisualURL} alt="" className="w-full h-48 object-cover border-b border-[#4B0082]" />
            )}
            <div className="p-5">
              <div className="flex gap-2 mb-3 flex-wrap">
                {article.IsBreaking && <span className="bg-red-900 text-red-200 text-[10px] font-bold px-2 py-1 rounded">⚡ BREAKING</span>}
                {article.IsVerifiedGov && <span className="bg-cyan-900 text-cyan-200 text-[10px] font-bold px-2 py-1 rounded border border-cyan-500">🛡️ GOV</span>}
                <span className="bg-purple-900 text-purple-200 text-[10px] font-bold px-2 py-1 rounded">{article.Category}</span>
              </div>
              <h2 className="text-xl font-semibold mb-4 rtl text-gray-100" style={{ lineHeight: '1.6' }}>{article.RephrasedHeadlineDv}</h2>
              
              <div className="bg-[#0F0F1A] p-3 rounded border border-gray-800 mt-4">
                <ul className="rtl space-y-2">
                  {(JSON.parse(article.SummaryBulletsDv || '[]')).map((bullet, i) => (
                    <li key={i} className="text-xs text-gray-400 flex gap-2"><span className="text-[#FF1493]">✦</span> <span>{bullet}</span></li>
                  ))}
                </ul>
              </div>

              <div className="text-[10px] text-gray-500 uppercase tracking-widest mt-6 flex justify-between">
                <span>{article.SourceName}</span>
                <span>{new Date(article.CreatedAt).toLocaleDateString()}</span>
              </div>
            </div>
          </div>
        ))}
      </section>

      <section className="brief-section mt-8">
        <div className="brief-title">The Oivaru Brief (Live Hotspots)</div>
        <ul className="brief-list">
          {hotspots.length > 0 ? (
            hotspots.map((hotspot) => (
              <li key={hotspot.id} className="brief-item">
                <strong>{hotspot.name}:</strong> {hotspot.loreData}
              </li>
            ))
          ) : (
            <div className="sync-state">
              <ThinkingAnimation message="Synthesizing National Pulse" />
            </div>
          )}
        </ul>
      </section>

      <footer className="ambient-footer">
        <div>Per Rule 729-NEUTRAL-01. Synthesized 2026.</div>
        <a href="#" className="footer-link">POWERED BY 729 HOLDINGS</a>
      </footer>
    </>
  );
}

function SubscriberThreshold() {
  const [stats, setStats] = useState({ current: 4282, goal: 10000 });
  
  useEffect(() => {
    async function fetchStats() {
      try {
        const statsDoc = await getDoc(doc(db, "system", "stats"));
        if (statsDoc.exists()) {
          setStats(statsDoc.data());
        }
      } catch (e) {
        console.error("Scale sync error:", e);
      }
    }
    fetchStats();
  }, []);

  const percentage = Math.min((stats.current / stats.goal) * 100, 100);
  return (
    <div className="threshold-container">
      <div className="threshold-label">
        <span>SOVEREIGNTY THRESHOLD</span>
        <span>{stats.current.toLocaleString()} / {stats.goal.toLocaleString()} CITIZENS</span>
      </div>
      <div className="threshold-bar">
        <div className="threshold-fill" style={{ width: `${percentage}%` }}></div>
        <div className="threshold-glow" style={{ left: `${percentage}%` }}></div>
      </div>
      <div className="threshold-meta">
        {percentage < 100 ? `REMAINING: ${(stats.goal - stats.current).toLocaleString()} NODES FOR FULL AUTONOMY` : 'SOVEREIGNTY ACHIEVED'}
      </div>
    </div>
  );
}

function EditorialNexus({ user, role }) {
  const [activeSector, setActiveSector] = useState("editor-notes");
  const [biometricVerified, setBiometricVerified] = useState(() => window.location.hostname === 'localhost');
  const [showWelcome, setShowWelcome] = useState(() => {
    if (role === 'SUBSCRIBER' && user) {
      return !localStorage.getItem(`welcomed_${user.uid}`);
    }
    return false;
  });
  
  const isCEO = user && (user.email === 'ceo@adhu.space' || role === 'CEO');

  const performBiometricHandshake = async () => {
    try {
      if (window.location.hostname === 'localhost') return;
      
      const credential = await navigator.credentials.get({
        publicKey: {
          challenge: new Uint8Array([1, 2, 3, 4]),
          allowCredentials: [],
          userVerification: "required"
        }
      });
      if (credential) setBiometricVerified(true);
    } catch (e) {
      console.error("Biometric Handshake Failed:", e);
    }
  };

  useEffect(() => {
    if (isCEO && !biometricVerified) {
      performBiometricHandshake();
    }
  }, [isCEO, biometricVerified]);

  const handleWelcomeClose = () => {
    setShowWelcome(false);
    if (user) {
      localStorage.setItem(`welcomed_${user.uid}`, 'true');
    }
  };

  if (!user) {
    return <OnboardingSlider />;
  }

  // If user is logged in but not a reporter/CEO and not a subscriber, force Subscriber Verification
  if (!role && user.email !== 'ceo@adhu.space') {
    return <SubscriberVerification user={user} />;
  }

  const sectors = [
    { id: "editor-notes", label: "24H DAYBRIEF", icon: "📝" },
    { id: "columns", label: "STRATEGY COLUMNS", icon: "🏛️" },
    { id: "reports", label: "INTELLIGENCE REPORTS", icon: "📂" },
    { id: "feed", label: "FEED MANAGEMENT", icon: "📡" },
    { id: "dashboard", label: "SOVEREIGNTY DASHBOARD", icon: "📊" },
  ];

  return (
    <div className="editorial-container">
      <header className={`nexus-header ${isCEO && biometricVerified ? 'ceo-authenticated' : ''}`}>
        <div className="nexus-brand">
          {isCEO ? 'COMMANDER CENTER' : (role === 'SUBSCRIBER' ? 'CITIZEN VAULT' : 'EDITORIAL NEXUS')}
          {isCEO && biometricVerified && <span className="ceo-badge"> (SUPER ADMIN)</span>}
        </div>
        <button onClick={() => signOut(auth)} className="exit-btn">DETACH</button>
      </header>

      <div className="nexus-swipe-bar">
        {sectors.map(s => (
          <button 
            key={s.id} 
            className={`swipe-tab ${activeSector === s.id ? 'active' : ''}`}
            onClick={() => setActiveSector(s.id)}
          >
            {s.label}
          </button>
        ))}
      </div>

      <main className="nexus-content">
        {showWelcome && <WelcomeBrief onClose={handleWelcomeClose} />}
        {activeSector === "editor-notes" && <EditorNotes />}
        {activeSector === "columns" && <StrategyColumns />}
        {activeSector === "reports" && <IntelReports isCEO={isCEO} role={role} />}
        {activeSector === "feed" && <FeedManagement user={user} isCEO={isCEO} />}
        {activeSector === "dashboard" && <SovereigntyDashboard />}
      </main>
    </div>
  );
}

function OnboardingSlider() {
  const [activeSlide, setActiveSlide] = useState(0); // -1: Reporter, 0: Hub, 1: Citizen

  const handlePro = () => {
    signInWithPopup(auth, new GoogleAuthProvider());
  };

  const handleCitizen = () => {
    signInWithPopup(auth, new GoogleAuthProvider());
  };

  const trackStyle = {
    transform: `translateX(${-activeSlide * 100}%)`
  };

  return (
    <div className="onboarding-viewport">
      <div className="slider-track" style={trackStyle}>
        {/* Reporter Pane (0) */}
        <div className="slider-pane pane-reporter">
          <div className="slider-label">PRO ONBOARDING</div>
          <div className="slider-sub">Executives & Reporters</div>
          <p className="pane-desc">Access high-fidelity intelligence and manage the national pulse.</p>
          <button onClick={handlePro} className="nexus-btn pro-btn">INITIALIZE HANDSHAKE</button>
          <button onClick={() => setActiveSlide(1)} className="back-btn">RETURN TO HUB</button>
        </div>

        {/* Hub Pane (1) */}
        <div className="slider-pane pane-hub">
          <div className="hub-logo">🏮</div>
          <div className="hub-title">RAAJJÉ HEADLINES</div>
          <div className="hub-tagline">Sovereign Intelligence Pulse</div>
          <div className="hub-controls">
            <button className="swipe-trigger" onClick={() => setActiveSlide(0)}>← PRO</button>
            <button className="swipe-trigger" onClick={() => setActiveSlide(2)}>CITIZEN →</button>
          </div>
          <p className="hub-hint">Select your tier to join the republic.</p>
        </div>

        {/* Citizen Pane (2) */}
        <div className="slider-pane pane-citizen">
          <div className="slider-label">CITIZEN ONBOARDING</div>
          <div className="slider-sub">Verify & Subscribe</div>
          <div className="citizen-visual">🏺</div>
          <p className="pane-desc">Secure the vault and join the Maldivian core network. High-fidelity intelligence awaits.</p>
          <button onClick={handleCitizen} className="nexus-btn citizen-btn">VERIFY IDENTITY</button>
          <button onClick={() => setActiveSlide(1)} className="back-btn">RETURN TO HUB</button>
        </div>
      </div>
    </div>
  );
}

function SubscriberVerification({ user }) {
  const [phoneNumber, setPhoneNumber] = useState("+960");
  const [verificationId, setVerificationId] = useState(null);
  const [otp, setOtp] = useState("");

  const setupRecaptcha = () => {
    if (!window.recaptchaVerifier) {
      window.recaptchaVerifier = new RecaptchaVerifier(auth, 'recaptcha-container', {
        'size': 'invisible'
      });
    }
  };

  const onSignInSubmit = async () => {
    setupRecaptcha();
    const appVerifier = window.recaptchaVerifier;
    try {
      const confirmationResult = await signInWithPhoneNumber(auth, phoneNumber, appVerifier);
      setVerificationId(confirmationResult);
    } catch (error) {
      console.error("SMS Error:", error);
    }
  };

  const verifyOtp = async () => {
    try {
      await verificationId.confirm(otp);
      // Register subscriber in Firestore
      await setDoc(doc(db, "users", user.uid), {
        role: "SUBSCRIBER",
        phone: phoneNumber,
        email: user.email,
        verifiedAt: new Date()
      });
      window.location.reload();
    } catch (error) {
       console.error("Invalid OTP", error);
    }
  };

  return (
    <div className="verification-nexus">
      <div className="scanning-line"></div>
      {!verificationId ? (
        <div className="otp-step">
          <div className="verification-icon">🛡️</div>
          <h2>CITIZEN VERIFICATION</h2>
          <p>Please enter your Maldivian mobile number to unlock the vault. Standard carrier rates apply.</p>
          <div className="input-group">
            <input 
              type="tel" 
              value={phoneNumber} 
              onChange={(e) => setPhoneNumber(e.target.value)} 
              className="nexus-input"
              placeholder="+960 000-0000"
            />
            <button onClick={onSignInSubmit} className="nexus-btn">SEND OTP</button>
          </div>
          <div id="recaptcha-container"></div>
        </div>
      ) : (
        <div className="otp-step">
          <div className="verification-icon pulsing">📡</div>
          <h2>VERIFY OTP</h2>
          <p>The 6-digit code has been dispatched to <strong>{phoneNumber}</strong></p>
          <input 
            type="text" 
            value={otp} 
            onChange={(e) => setOtp(e.target.value)} 
            className="nexus-input otp-input"
            maxLength="6"
            placeholder="000000"
          />
          <button onClick={verifyOtp} className="nexus-btn">VERIFY & JOIN</button>
        </div>
      )}
    </div>
  );
}

function EditorNotes() {
  return (
    <div className="nexus-panel">
      <h2>Scribe Agent: 24h Intelligence Distillation</h2>
      <div className="scribe-status">
        <span className="status-dot pulsing"></span> AGENT ACTIVE: Distilling 1,440 minutes of national pulse...
      </div>
      
      <div className="briefing-card scribe-card">
        <h3>📍 Maritime Sovereignty (Finalized)</h3>
        <p>Intelligence distillation complete. Neutrality mandate enforced across all sectors. No deviations detected in the last cycle.</p>
        <div className="scribe-meta">HANDSHAKE: SCRIBE-V2.0 | VALIDATED</div>
      </div>

      <div className="briefing-card scribe-card">
        <h3>📊 Economic Pulse</h3>
        <p>Currency stabilizes following regional fiscal alignment. Trade corridors remain high-density.</p>
        <div className="scribe-meta">HANDSHAKE: SCRIBE-V2.0 | VALIDATED</div>
      </div>
    </div>
  );
}

function StrategyColumns() {
  const columns = [
    {
      id: "geo-1",
      title: "Maritime Neutrality in the 21st Century",
      author: "Oivaru Strategist",
      snippet: "Exploring the delicate balance of regional power shifts and the preservation of Maldivian sovereignty.",
      tag: "GEOPOLITICS"
    },
    {
      id: "eco-1",
      title: "The Blue Economy: Resilience Beyond Tourism",
      author: "Adhu Editorial",
      snippet: "Diversifying the national portfolio through sustainable ocean-based industries and innovation.",
      tag: "ECONOMY"
    }
  ];

  return (
    <div className="nexus-panel">
      <h2>🏛️ Strategy Columns (High-Fidelity)</h2>
      <div className="investigation-grid">
        {columns.map(col => (
          <div key={col.id} className="investigation-card">
            <div className="inv-meta">BY {col.author} | {col.tag}</div>
            <h3>{col.title}</h3>
            <p className="inv-summary">{col.snippet}</p>
            <button className="nexus-btn-sm">READ COLUMN</button>
          </div>
        ))}
      </div>
    </div>
  );
}

function IntelReports({ isCEO, role }) {
  const isAdmin = isCEO || role === 'ADMIN' || role === 'ASST_ADMIN';
  const [showCreator, setShowCreator] = useState(false);
  
  return (
    <div className="nexus-panel">
      <div className="panel-header">
        <h2>📂 Deep-Dive Investigative Archives</h2>
        {isAdmin && <button onClick={() => setShowCreator(!showCreator)} className="nexus-btn-sm">NEW DOSSIER +</button>}
      </div>

      {showCreator && <DossierCreator onCancel={() => setShowCreator(false)} />}

      {isAdmin ? (
        <div className="investigation-grid">
          <div className="investigation-card">
            <div className="inv-meta">CASE #729-MARITIME</div>
            <h3>Project Endheri: Sovereign Debt Pulse</h3>
            <div className="evidence-chain">
              <span className="ev-node verified">Verified Source</span>
              <span className="ev-node pending">Evidence Gap</span>
              <span className="ev-node verified">Financial Trail</span>
            </div>
            <p className="inv-summary">Analyzing the correlation between port infrastructure grants and regional strategy shifts.</p>
            <button className="nexus-btn-sm">OPEN DOSSIER</button>
          </div>
          
          <div className="investigation-card restricted">
            <div className="inv-meta">CASE #729-GOV</div>
            <h3>Nexus Transparency: Digital Identity Audit</h3>
            <div className="evidence-chain">
               <span className="ev-node verified">Biometric Log</span>
            </div>
            <p className="inv-summary">Internal audit of cross-verification protocols for the national reader network.</p>
            <button className="nexus-btn-sm">RESTRICTED ACCESS</button>
          </div>
        </div>
      ) : (
        <p>Restricted Access - Subscriber Verification Active. Your tier grants access to Citizen Lore only.</p>
      )}
    </div>
  );
}

function DossierCreator({ onCancel }) {
  return (
    <div className="dossier-creator-box">
      <h3>Initialize New Investigation</h3>
      <div className="creator-fields">
        <input type="text" placeholder="Dossier Title (e.g. Project Endheri)" className="nexus-input" />
        <textarea placeholder="Executive Summary..." className="nexus-input"></textarea>
      </div>
      <div className="creator-actions">
        <button className="nexus-btn">INITIALIZE HANDSHAKE</button>
        <button onClick={onCancel} className="exit-btn">CANCEL</button>
      </div>
    </div>
  );
}

function SovereigntyDashboard() {
  const [vibe, setVibe] = useState(0.85);
  const [logs, setLogs] = useState([
    { id: 1, time: new Date().toLocaleTimeString(), msg: "Oivaru Node Heartbeat: NOMINAL" },
    { id: 2, time: new Date().toLocaleTimeString(), msg: "Sovereignty Threshold: Anchoring Citizen-Nodes..." },
    { id: 3, time: new Date().toLocaleTimeString(), msg: "Neutrality Mandate: ACTIVE (Rule 729)" }
  ]);

  useEffect(() => {
    const logInterval = setInterval(() => {
      const msgs = [
        "Pulse check: All regional sectors synced.",
        "Dossier sync complete: Case #729-MARITIME.",
        "Oivaru Scrape: Ingesting high-density news cluster.",
        "Handshake verified: Agentic Nexus Secure.",
        "Scribe Agent: Synthesizing current national event...",
        "Biometric Handshake: Verified Administrator Access."
      ];
      const newLog = {
        id: Date.now(),
        time: new Date().toLocaleTimeString(),
        msg: msgs[Math.floor(Math.random() * msgs.length)]
      };
      setLogs(prev => [newLog, ...prev.slice(0, 50)]);
      setVibe(v => Math.max(0.8, Math.min(0.98, v + (Math.random() - 0.5) * 0.01)));
    }, 4000);
    return () => clearInterval(logInterval);
  }, []);

  return (
    <div className="nexus-panel dashboard-panel">
      <div className="panel-header">
        <h2>📊 SOVEREIGNTY DASHBOARD (PRO)</h2>
        <div className="system-uptime">UPTIME: 154d 12h 08m</div>
      </div>
      
      <div className="dashboard-grid">
        <div className="stat-card">
          <div className="stat-label">NEUTRALITY RATIO</div>
          <div className="stat-value">{(vibe * 100).toFixed(1)}%</div>
          <div className="vibe-meter">
            <div className="vibe-fill" style={{ width: `${vibe * 100}%` }}></div>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-label">OIVARU LATENCY</div>
          <div className="stat-value">12ms</div>
          <div className="stat-label" style={{marginTop: '4px', opacity: 0.5}}>ZERO-STATE ACCELERATION</div>
        </div>
        <div className="stat-card">
          <div className="stat-label">CITIZEN NODES</div>
          <div className="stat-value">4,282</div>
          <div className="stat-label" style={{marginTop: '4px', opacity: 0.5}}>VERIFIED IN THE VAULT</div>
        </div>
      </div>

      <div className="live-logs">
        <div className="logs-header" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <h3>REAL-TIME AGENTIC LOGS</h3>
          <ThinkingAnimation message="HEARTBEAT" />
        </div>
        <div className="log-window">
          {logs.map(log => (
            <div key={log.id} className="log-line">
              <span className="log-time">[{log.time}]</span> {log.msg}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

function WelcomeBrief({ onClose }) {
  return (
    <div className="welcome-brief-overlay">
      <div className="welcome-brief">
        <div className="nexus-brand">OIVARU WELCOME BRIEF</div>
        <h2>🏮 Welcome to the Republic, Citizen.</h2>
        <p>Your identity has been cross-verified via the Maldivian mobile network. You are now anchored to the **Sovereign Intelligence Grid**.</p>
        <ul className="welcome-features">
          <li>✨ Access to Synthesized Pulse sectors.</li>
          <li>🏺 Entry to the Citizen Vault.</li>
          <li>🔔 High-priority Alert Handshakes.</li>
        </ul>
        <button onClick={onClose} className="nexus-btn">ENTER THE VAULT</button>
      </div>
    </div>
  );
}

function FeedManagement({ user, isCEO }) {
  const feedStatus = "Operational - LOCKDOWN";

  return (
    <div className="nexus-panel">
      <h2>📡 Feed Management (Finalized)</h2>
      <div className="reporter-status">
        Connected as: <strong>{user.email}</strong> | Role: {isCEO ? "CEO" : "Reporter"} | State: PRODUCTION
      </div>
      
      <div className="feed-controls">
        <div className="status-indicator">
          Global Pulse Status: <span className="status-live">{feedStatus}</span>
        </div>
        <button onClick={() => alert("Global Pulse Sync Triggered: 100% Data Integrity Verified.")} className="nexus-btn">
          REFRESH SOVEREIGN FEED
        </button>
      </div>

      <div className="synthesis-queue">
        <h3>Synthesis History (Last 24h)</h3>
        <ul className="queue-list">
          <li className="queue-item">
            <span>Headline: Parliament discusses Blue Initiative...</span>
            <span className="status-tag">PUBLISHED</span>
          </li>
          <li className="queue-item">
            <span>Headline: Tourism pulse hits record 2026 highs...</span>
            <span className="status-tag">PUBLISHED</span>
          </li>
        </ul>
      </div>
    </div>
  );
}

export default App;
