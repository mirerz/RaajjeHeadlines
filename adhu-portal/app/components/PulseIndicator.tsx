import React from 'react';

const PulseIndicator = ({ intensity = 0.8, color = "#00ffff" }) => {
  return (
    <div className="flex flex-col items-center space-y-4">
      <div className="relative w-16 h-16">
        {/* Glow effect */}
        <div 
          className="absolute inset-0 rounded-full animate-pulse blur-xl opacity-50"
          style={{ backgroundColor: color }}
        ></div>
        {/* Core */}
        <div 
          className="absolute inset-2 rounded-full border-2 shadow-[0_0_20px_rgba(0,255,255,0.8)]"
          style={{ borderColor: color, backgroundColor: 'rgba(0,255,255,0.1)' }}
        ></div>
      </div>
      <div className="text-center">
        <span className="text-[10px] font-bold tracking-[0.2em] text-cyan-400 uppercase">Sentinel Stem</span>
        <div className="text-[14px] font-medium text-white tracking-widest">{Math.round(intensity * 100)}% LUME</div>
      </div>
    </div>
  );
};

export default PulseIndicator;
