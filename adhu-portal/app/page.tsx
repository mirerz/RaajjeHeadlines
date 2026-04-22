'use client';

import { useEffect, useState } from 'react';
import Image from 'next/image';
import SentinelPulse from './components/SentinelPulse';

export default function Page() {
  const [feed, setFeed] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function loadFeed() {
      // Logic: Use the public production API if available, otherwise fallback to local sandbox
      const baseUrl = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:3005";
      try {
        const res = await fetch(`${baseUrl}/api/feed`);
        if (res.ok) {
          const data = await res.json();
          setFeed(data);
        }
      } catch (e) {
        console.error("adhu.space API Handshake Failed:", e);
      } finally {
        setLoading(false);
      }
    }
    loadFeed();
  }, []);

  return (
    <main className="sterile-grid">
      <header className="flex justify-between items-center mb-8 pt-4">
        <Image src="/logo.png" alt="Raajjé Headlines" width={140} height={46} />
        <SentinelPulse />
      </header>
      
      {loading ? (
        <div className="glass-card flex items-center justify-center h-64">
          <p className="animate-pulse tracking-[0.2em] text-cyan-400">INITIALIZING INTELLIGENCE GRID...</p>
        </div>
      ) : feed.length > 0 ? (
        feed.map((article, index) => (
          <div key={article.ID || index} className="glass-card mb-6 border-l-2 border-cyan-500/30">
            <div className="flex gap-4 mb-2 items-center">
              <span className="text-[10px] font-mono text-cyan-400 tracking-widest uppercase">
                {article.Category || 'Intelligence'} Node
              </span>
              {article.IsVerifiedGov && (
                <span className="text-[8px] bg-cyan-900/50 text-cyan-200 px-2 py-0.5 rounded border border-cyan-500/20">
                  🛡️ VERIFIED GOV
                </span>
              )}
            </div>
            <h2 className="text-xl font-bold mb-3 rtl leading-relaxed">{article.RephrasedHeadlineDv}</h2>
            
            {/* Intelligence Briefing (Parsed JSON bullets) */}
            <div className="mt-4 space-y-2 bg-black/20 p-4 rounded">
               {(JSON.parse(article.SummaryBulletsDv || '[]')).map((bullet: string, i: number) => (
                  <div key={i} className="flex gap-2 items-start rtl">
                    <span className="text-cyan-500 mt-1">✦</span>
                    <p className="text-sm text-gray-300 leading-relaxed">{bullet}</p>
                  </div>
               ))}
            </div>

            <div className="mt-6 flex justify-between items-center">
              <span className="text-[9px] font-mono text-gray-500 uppercase tracking-tighter">
                SOURCE: {article.SourceName} | TIER {article.SourceTier}
              </span>
              <button className="text-[10px] text-cyan-400 hover:text-white transition uppercase tracking-widest font-bold">
                Access Node →
              </button>
            </div>
          </div>
        ))
      ) : (
        <div className="glass-card h-64 flex items-center justify-center">
          <p className="text-gray-500">NO ACTIVE SIGNALS DETECTED. GRID NOMINAL.</p>
        </div>
      )}

      <footer className="footer-pulse">
        [PULSE: GDP Growth 1.0%] | Raajjé HEADLINES | 729-NEUTRAL-01
      </footer>
    </main>
  );
}
