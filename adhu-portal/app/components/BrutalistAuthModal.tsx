'use client';

import React from 'react';

export default function BrutalistAuthModal({ isOpen, onClose }: { isOpen: boolean, onClose: () => void }) {
  if (!isOpen) return null;

  const socialNodes = [
    { name: 'Google', color: 'bg-white text-black border-black' },
    { name: 'X / Twitter', color: 'bg-black text-white border-white/20' },
    { name: 'Facebook', color: 'bg-[#1877F2] text-white border-white/10' },
    { name: 'Instagram', color: 'bg-gradient-to-tr from-[#f9ce34] via-[#ee2a7b] to-[#6228d7] text-white border-none' },
    { name: 'TikTok', color: 'bg-black text-white border-cyan-400 border-r-2 border-b-2' },
    { name: 'Microsoft', color: 'bg-[#00a1f1] text-white border-none' },
  ];

  return (
    <div className="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/80 backdrop-blur-md">
      <div className="bg-[#f4f1ea] border-[4px] border-black p-8 w-full max-w-md shadow-[12px_12px_0px_0px_rgba(0,0,0,1)] relative">
        <button onClick={onClose} className="absolute top-4 right-4 text-2xl font-black hover:scale-125 transition-transform">✕</button>
        
        <h2 className="text-3xl font-black uppercase tracking-tighter mb-2 italic">Join the Republic</h2>
        <p className="text-[10px] uppercase tracking-widest mb-8 text-black/60 font-bold border-b border-black/10 pb-4">
          Authenticate with your Sovereign Node
        </p>

        <div className="grid grid-cols-1 gap-3">
          {socialNodes.map((node) => (
            <button 
              key={node.name}
              className={`w-full p-4 font-black uppercase tracking-tighter text-sm flex justify-between items-center border-2 transition-all hover:-translate-y-1 hover:translate-x-1 hover:shadow-[-4px_4px_0px_0px_rgba(0,0,0,1)] ${node.color}`}
            >
              <span>{node.name}</span>
              <span className="text-[10px]">Connect →</span>
            </button>
          ))}
        </div>

        <div className="mt-8 text-[9px] text-center uppercase tracking-tighter font-bold text-black/40">
          By connecting, you agree to the 729 Holdings Neutrality Pact.
        </div>
      </div>
    </div>
  );
}
