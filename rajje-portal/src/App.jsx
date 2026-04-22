import './index.css';

function App() {
  return (
    <>
      <header className="ambient-header">
        <div className="logo">
          adhu.space<span className="dot">.</span>
        </div>
        <div className="vibe-indicator">
          <div className="pulse-dot"></div>
          Cabinet Vibe: Yellow (Logistics Review)
        </div>
      </header>

      <section className="hero-section">
        <h1 className="hero-title">
          Global Rights Coalition Urges Immediate Halt to Death Penalty Resurgence
        </h1>
        <div 
          className="oivaru-tag"
          onClick={() => alert('Primary Sources:\n- Presidency MV (Official Press)\n- Maldives Independent (Op-Ed)\n- Human Rights Watch (Data)\n\nNote: Sterilized per 729-NEUTRAL-01 mandate.')}
        >
          <span>✨</span> Synthesized by Oivaru
        </div>
      </section>

      <section className="matrix-section">
        <div className="matrix-title">
          Social Vibe Density Matrix
        </div>
        <svg width="400" height="400" viewBox="0 0 100 200" xmlns="http://www.w3.org/2000/svg">
            <circle cx="50" cy="20" r="4" className="atoll" />
            <circle cx="55" cy="40" r="5" className="atoll" />
            <circle cx="45" cy="60" r="7" className="atoll" style={{ fill: 'rgba(255, 215, 0, 0.4)', stroke: '#FFD700' }} /> 
            <circle cx="60" cy="80" r="6" className="atoll" />
            <circle cx="50" cy="100" r="8" className="atoll matrix-glow" /> 
            <circle cx="40" cy="120" r="5" className="atoll" />
            <circle cx="65" cy="140" r="4" className="atoll" />
            <circle cx="50" cy="170" r="5" className="atoll" style={{ fill: 'rgba(255, 20, 147, 0.4)', stroke: '#FFFF1493' }} /> 
            <path d="M50 20 L55 40 L45 60 L60 80 L50 100 L40 120 L50 170" stroke="rgba(255,255,255,0.05)" strokeWidth="1" fill="none" />
        </svg>
      </section>

      <section className="brief-section">
        <div className="brief-title">
          The Oivaru Brief (30 Sec)
        </div>
        <ul className="brief-list">
          <li className="brief-item">
            The 15-member cabinet shift is stabilizing, with public sentiment migrating from partisan anger to logistical policy analysis.
          </li>
          <li className="brief-item">
            Energy import constraints tied to broader Middle East unrest remain the driving force behind the ADB's cautious 1.0% growth trajectory.
          </li>
          <li className="brief-item">
            State infrastructure projects are signaling shifts geared heavily toward "Independence and Resilience," moving away from legacy expansion strategies.
          </li>
        </ul>
      </section>

      <footer className="ambient-footer">
        <div>Per Rule 729-NEUTRAL-01. Synthesized 2026.</div>
        <a href="#" className="footer-link">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" className="footer-icon">
              <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path>
              <polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline>
              <line x1="12" y1="22.08" x2="12" y2="12"></line>
          </svg>
          POWERED BY ENDHERI ODI
        </a>
      </footer>
    </>
  );
}

export default App;
