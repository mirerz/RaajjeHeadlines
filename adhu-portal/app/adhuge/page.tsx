'use client';
import { useState } from 'react';
import Image from 'next/image';
import SentinelPulse from '../components/SentinelPulse';
import CabinetGrid from '../components/CabinetGrid';
import PulseIndicator from '../components/PulseIndicator';

const THAANA_CATEGORIES = [
  { eng: "National", tha: "ޚަބަރު" },
  { eng: "Environment", tha: "ތިމާވެށި" },
  { eng: "Economy", tha: "އިޤްތިޞާދު" },
  { eng: "Politics", tha: "ސިޔާސަތު" },
  { eng: "Youth & Women", tha: "ޒުވާނުންނާއި އަންހެނުން" },
  { eng: "Business", tha: "ވިޔަފާރި" },
  { eng: "Fashion", tha: "ފެޝަން" },
  { eng: "Kids", tha: "ކުޑަކުދިން" },
];

export default function AdhugePortal() {
  const [activeCategory, setActiveCategory] = useState("ޚަބަރު");
  const [seniorMode, setSeniorMode] = useState(false);
  const [neutrality, setNeutrality] = useState(true);

  // Toggle colors based on accessibility mode
  const bgStyle = seniorMode ? { backgroundColor: '#F8FAFC', color: '#0B0E23' } : {};
  const cardStyle = seniorMode ? { background: '#FFF', border: '1px solid #CBD5E1', color: '#0B0E23', boxShadow: '0 4px 6px -1px rgb(0 0 0 / 0.1)' } : {};
  const textScale = seniorMode ? 'text-xl' : 'text-base';
  const headerScale = seniorMode ? 'text-3xl' : 'text-xl';

  return (
    <main className="sterile-grid" style={{ ...bgStyle, minHeight: '100vh', paddingBottom: '4rem' }}>
      {/* Header Container */}
      <header className="flex justify-between items-center mb-4">
        <div className="flex-1">
          <div className="mb-2">
            <Image src="/logo.png" alt="Raajjé Headlines" width={120} height={40} className="opacity-90" />
          </div>
          <div className="text-[10px] opacity-60 font-bold mb-1 tracking-widest">OIVARU ENGINE {neutrality ? '[STERILE]' : '[OVERRIDE]'}</div>
          {/* Senior Accessibility Toggle */}
          <button 
            onClick={() => setSeniorMode(!seniorMode)} 
            className="thaana-pill text-xs px-3 py-1"
            style={seniorMode ? { background: '#0B0E23', color: '#FFF' } : {}}
          >
            {seniorMode ? 'STANDARD MODE' : 'SENIOR MODE (HIGH CONTRAST)'}
          </button>
        </div>
        <div className="flex-1 flex justify-center flex-col items-center">
          <SentinelPulse />
          <h1 className={`kinetic-thaana font-bold mt-2 ${headerScale}`} style={{ color: seniorMode ? '#0B0E23' : 'var(--bioluminescent-cyan)' }}>އަދުގޭ</h1>
        </div>
        <div className="flex-1 flex flex-col items-end gap-2 pr-2">
          <div className="text-right text-[10px] opacity-80 uppercase tracking-tighter">
            <div>29°C Male'</div>
            <div>Asr: 15:24</div>
          </div>
          <PulseIndicator intensity={0.92} color="#00ffff" />
        </div>
      </header>

      {/* Horizontal Nav */}
      <nav className="horizontal-nav pb-2 border-b border-white/10" style={seniorMode ? { borderColor: '#CBD5E1' } : {}}>
        {THAANA_CATEGORIES.map(cat => (
          <div 
            key={cat.tha} 
            className={`thaana-pill ${activeCategory === cat.tha ? 'active' : ''}`}
            onClick={() => setActiveCategory(cat.tha)}
            style={seniorMode && activeCategory !== cat.tha ? { background: '#E2E8F0', color: '#0F172A', borderColor: '#CBD5E1' } : {}}
          >
            <div className="text-lg text-center font-bold">{cat.tha}</div>
            <div className="text-[0.65rem] opacity-80 text-center uppercase tracking-wider">{cat.eng}</div>
          </div>
        ))}
      </nav>

      {/* Stakeholder Vibe Matrix */}
      <section className="flex flex-col gap-2">
        <div className="flex justify-between items-center px-1">
          <span className="text-[0.6rem] font-bold tracking-[0.2em] opacity-50">STAKEHOLDER VIBE MATRIX</span>
          <span className="text-[0.6rem] font-bold text-cyan-400">LIVE FEED</span>
        </div>
        <div className="flex gap-2 h-1 overflow-hidden">
          <div className="flex-1 bg-cyan-400/50 animate-pulse"></div>
          <div className="flex-1 bg-yellow-400/50"></div>
          <div className="flex-1 bg-white/20"></div>
        </div>
        <div className="flex justify-between text-[0.5rem] opacity-40 uppercase">
          <span>Institutions</span>
          <span>Seniors</span>
          <span>Influencers</span>
        </div>
      </section>

      {/* Cabinet Portfolios */}
      <section>
        <h3 className="text-xs font-bold opacity-70 mb-2 px-1">GOVERNANCE PULSE: 15-MEMBER CABINET</h3>
        <CabinetGrid />
      </section>

      {/* The Live News Grids */}
      <section className="flex flex-col gap-6">

        {/* Politics / High Conflict - REFACTORED FOR NEUTRALITY */}
        <div className="glass-card flex flex-col gap-3" style={cardStyle}>
          <div className="flex justify-between items-center text-sm opacity-70">
            <span className="uppercase tracking-widest text-[0.6rem] font-bold">ސިޔާސަތު (Political Analysis)</span>
            <div className="flex items-center gap-2">
              <span className="font-medium text-[0.6rem] uppercase" style={{ color: seniorMode ? '#000' : 'var(--bioluminescent-cyan)' }}>729-NEUTRAL-01 VERIFIED</span>
              <div className="pulse-dot" style={{ backgroundColor: 'var(--bioluminescent-cyan)', width: '6px', height: '6px' }}></div>
            </div>
          </div>
          <h2 className={`kinetic-thaana font-bold leading-tight ${seniorMode ? 'text-2xl' : 'text-lg'}`} style={{ color: seniorMode ? '#0B0E23' : '#FFF' }}>
            މަރުގެ އަދަބާ ގުޅޭ ބިލު: ކަމާބެހޭ ފަރާތްތަކުން މަޝްވަރާ ފަށައިފި
          </h2>
          <p className="sterile-text text-xs opacity-70">
            The updated draft seeks to align administrative protocols with existing legislative frameworks. 
            International observers are monitoring the standard of procedure.
          </p>
        </div>

        {/* Economy / Stable */}
        <div className="glass-card flex flex-col gap-3" style={cardStyle}>
          <div className="flex justify-between items-center text-sm opacity-70">
            <span className="uppercase tracking-widest text-[0.6rem] font-bold">އިޤްޞާދު (Fiscal Data)</span>
          </div>
          <h2 className={`kinetic-thaana font-bold leading-tight ${seniorMode ? 'text-2xl' : 'text-lg'}`} style={{ color: seniorMode ? '#0B0E23' : '#FFF' }}>
            އޭޑީބީގެ ދިރާސާ: މިއަހަރުގެ އިޤްތިޞާދީ އަންދާޒާ ޕަބްލިޝްކޮށްފި
          </h2>
        </div>

        {/* National / Stable */}
        <div className="glass-card flex flex-col gap-3" style={cardStyle}>
          <div className="flex justify-between items-center text-sm opacity-70">
            <span className="uppercase tracking-widest text-[0.65rem] font-bold">ޚަބަރު (National)</span>
            <div className="flex items-center gap-2">
              <span className="font-medium" style={{ color: '#32CD32' }}>Stable</span>
              <div className="pulse-dot" style={{ backgroundColor: '#32CD32', animationDuration: '3s' }}></div>
            </div>
          </div>
          <h2 className={`kinetic-thaana font-bold leading-tight ${seniorMode ? 'text-4xl' : 'text-xl'}`} style={{ color: seniorMode ? '#0B0E23' : '#FFF' }}>
            ސަރުކާރުގެ އައު އޮނިގަނޑު: 15 މިނިސްޓްރީއާއެކު ހިންގުން ހަރުދަނާކުރުމުގެ މަސައްކަތް ފަށައިފި
          </h2>
        </div>

      </section>

      {/* Endheri Odi Heritage Footer */}
      <section className="glass-card flex justify-between items-center bg-cyan-900/20" style={seniorMode ? { background: '#F0F9FF', border: '1px solid #BAE6FD' } : {}}>
        <div>
          <h3 className="font-bold text-lg" style={{ color: 'var(--bioluminescent-cyan)' }}>Endheri Odi Heritage</h3>
          <p className="text-xs opacity-80 mt-1">Tap to launch 3D Cultural Anchor & Historical Shipwright Data.</p>
        </div>
        <div className="h-10 w-10 rounded-full border-2 border-cyan-400 flex items-center justify-center animate-pulse cursor-pointer shadow-[0_0_15px_rgba(0,255,255,0.4)]">
          <span className="text-xl">⛵</span>
        </div>
      </section>

    </main>
  );
}
