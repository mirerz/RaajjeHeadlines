'use client';

import { useEffect, useState } from 'react';
import Image from 'next/image';
import BrutalistAuthModal from './components/BrutalistAuthModal';

export default function Page() {
  const [feed, setFeed] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [isAuthOpen, setIsAuthOpen] = useState(false);

  useEffect(() => {
    async function loadFeed() {
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

  const [expandedCategory, setExpandedCategory] = useState<string | null>(null);

  return (
    <main className="broadsheet-container">
      <BrutalistAuthModal isOpen={isAuthOpen} onClose={() => setIsAuthOpen(false)} />
      
      {/* Full-Page Category Environment (Expansion State) */}
      {expandedCategory && (
        <div className="fixed inset-0 z-[200] bg-midnight-heritage p-12 animate-in fade-in zoom-in duration-500">
           <button onClick={() => setExpandedCategory(null)} className="absolute top-8 right-12 text-prestige-gold text-4xl font-black">✕</button>
           <h1 className="text-prestige-gold text-7xl font-black uppercase tracking-tighter mb-8 italic">{expandedCategory} Environment</h1>
           <div className="grid grid-cols-3 gap-12">
              {[1, 2, 3].map(i => <div key={i} className="glass-card h-96 border-prestige-gold/20"></div>)}
           </div>
        </div>
      )}

      {/* 1. Vertical Sidebar (Category Icons) */}
      <aside className="flex flex-col gap-8 items-center pt-8 border-r border-black/10 h-screen sticky top-0">
         {[
           { icon: '🏛️', label: 'Politics' },
           { icon: '🌍', label: 'World' },
           { icon: '📈', label: 'Economy' },
           { icon: '⚖️', label: 'Justice' },
           { icon: '🎨', label: 'Culture' }
         ].map((cat, i) => (
           <div 
             key={i} 
             onMouseDown={() => {
                const timer = setTimeout(() => setExpandedCategory(cat.label), 800);
                (window as any)._catTimer = timer;
             }}
             onMouseUp={() => clearTimeout((window as any)._catTimer)}
             className="cursor-pointer hover:scale-125 transition-transform text-2xl filter grayscale hover:grayscale-0 relative group"
           >
             {cat.icon}
             <span className="absolute left-full ml-4 px-2 py-1 bg-black text-white text-[10px] uppercase opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">Hold to Expand</span>
           </div>
         ))}
      </aside>

      {/* 2. Main Broadsheet Feed */}
      <section className="main-content">
        <header className="mb-12 flex flex-col items-center">
          <div className="w-full border-t border-black mb-4"></div>
          <div className="py-4 w-full flex justify-center">
            <Image src="/logo-deepblue.png" alt="Raajjé HEADLINES" width={700} height={180} className="object-contain" priority />
          </div>
          <div className="w-full border-y-2 border-black py-2 mb-8 flex justify-between px-4 text-[10px] font-black uppercase tracking-[0.4em] font-mono text-black">
            <span>VOL. I ... NO. 01</span>
            <span className="text-red-600">MALDIVES SOVEREIGN GRID</span>
            <span>WEDNESDAY, APRIL 22, 2026</span>
          </div>
        </header>

        {loading ? (
          <div className="animate-pulse flex flex-col items-center justify-center h-64">
            <div className="pulse-hex"></div>
            <p className="text-xs tracking-widest uppercase mt-4">Synthesizing Front Page...</p>
          </div>
        ) : (
          <div className="flex flex-col gap-12 sticky-crawl">
            {feed.map((article, index) => (
              <article key={article.ID || index} className={`p-8 in-view relative group overflow-hidden transition-all duration-700 ${index === 0 ? 'news-item-high-rank' : 'border-b border-black/5 pb-12'}`}>
                {/* 3D Page Turn Hint */}
                <div className="absolute top-0 right-0 w-8 h-8 bg-black/5 border-l border-b border-black/10 group-hover:w-full group-hover:h-full group-hover:bg-broadsheet-cream/90 transition-all duration-700 z-10 flex items-start justify-end p-2 opacity-0 group-hover:opacity-100 cursor-pointer">
                   <span className="text-[10px] font-black uppercase tracking-tighter">Turn Page →</span>
                </div>

                {index === 0 && <span className="text-[10px] bg-black text-white px-2 py-1 font-bold mb-4 inline-block uppercase tracking-widest">Top Intelligence Node</span>}
                <h2 className={`${index === 0 ? 'hero-headline' : 'text-3xl font-bold'} rtl-broadsheet mb-6`}>
                  {article.RephrasedHeadlineDv}
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-8 rtl-broadsheet">
                   <p className="text-lg leading-relaxed text-black/80">{article.RephrasedBodyDv?.substring(0, 300)}...</p>
                   <div className="bg-black/5 p-6 border-r-4 border-[#F4B400]">
                      <h4 className="text-xs font-bold uppercase mb-4 tracking-widest">Intelligence Briefing</h4>
                      <ul className="space-y-3">
                        {(JSON.parse(article.SummaryBulletsDv || '[]')).map((bullet: string, i: number) => (
                          <li key={i} className="text-sm">✦ {bullet}</li>
                        ))}
                      </ul>
                   </div>
                </div>
                <div className="mt-8 pt-4 border-t border-black/5 flex justify-between items-center text-[10px] font-bold uppercase tracking-widest text-black/40">
                   <span>Source: {article.SourceName}</span>
                   <button className="hover:text-black transition">Read Full Broad-Sheet →</button>
                </div>
              </article>
            ))}
          </div>
        )}
      </section>

      {/* 3. The 5-Tile Transformer Grid */}
      <aside className="transformer-grid">
         <div className="tile-1x1 border-red-500 text-red-500 bg-red-500/5">Breaking</div>
         <div className="tile-1x1">Political</div>
         <div className="tile-1x1 border-cyan-500 text-cyan-500" onClick={() => setIsAuthOpen(true)}>Adhuge</div>
         <div className="tile-1x1">Dropick</div>
         <div className="tile-1x1 bg-[#F4B400] text-[#000B1E]" onClick={() => setIsAuthOpen(true)}>Profile</div>
      </aside>

      {/* 4. Portal Core (AI Assistant FAB) */}
      <div 
        onClick={() => setIsAuthOpen(true)}
        className="fixed bottom-8 right-8 w-16 h-16 bg-[#000B1E] rounded-full flex items-center justify-center shadow-2xl border-2 border-[#F4B400] cursor-pointer hover:rotate-90 transition-transform group"
      >
          <div className="w-4 h-4 bg-[#F4B400] rounded-sm group-hover:scale-150 transition-transform animate-pulse"></div>
      </div>
    </main>
  );
}

