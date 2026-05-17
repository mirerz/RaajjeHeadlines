import Image from 'next/image';

export default function Maadhaa27Page() {
  return (
    <div className="min-h-screen bg-[#E5E5E5] pb-20 font-sans">
      
      {/* HEADER SECTION */}
      <header className="bg-[#050A24] w-full flex flex-col items-center pt-8 pb-0">
        {/* adhu space logo */}
        <div className="text-6xl font-light tracking-widest text-[#00AEEF] flex items-center justify-center gap-1 mb-2">
          adhu
          <span className="text-red-600 font-medium relative">
             <div className="absolute -left-3 top-1/2 -translate-y-1/2 w-2 h-2 rounded-full border-2 border-red-600"></div>
             space
          </span>
        </div>
        
        {/* Top Info Bar */}
        <div className="w-full flex justify-between items-center text-white text-[10px] border-y border-white/20 mt-4 px-8 py-2 max-w-[1400px]">
          <div className="flex gap-4">
             {/* Left side */}
          </div>
          <div className="flex gap-8 rtl" dir="rtl">
            <span>1447 ރަބީޢުލްއައްވަލް 14 | 15 މެއި 2026</span>
            <span>މާލެ - ރާއްޖެ</span>
            <span>ހޫނުމިން: 32°C</span>
          </div>
        </div>
      </header>

      {/* MAIN CONTENT */}
      <main className="w-full max-w-[1400px] mx-auto mt-8 flex flex-col gap-10 px-8 relative">
        
        {/* GREEN BELT (Live Player Placeholder) */}
        <section className="relative w-full">
          {/* Note text left of the container */}
          <div className="absolute -left-48 top-1/2 -translate-y-1/2 text-red-600 text-[10px] w-40 text-right font-bold leading-tight z-0">
             Note:<br/>
             Clubhouse - Live 22:30<br/>
             (except Saturday Nigh)<br/>
             Clubhouse Feed Here...........{'>'}
          </div>

          <div className="flex w-full h-[140px] shadow-xl relative z-10">
            {/* Dark Green Belt */}
            <div className="flex-1 bg-[#092B14] border-y-2 border-l-2 border-gray-600 rounded-l-[70px] flex items-center justify-between px-10 relative overflow-hidden">
               {/* Left lines */}
               <div className="absolute left-0 top-6 w-20 border-t border-b border-yellow-600/50 h-2"></div>
               <div className="absolute left-0 bottom-6 w-20 border-t border-b border-yellow-600/50 h-2"></div>

               {/* Clubhouse Section */}
               <div className="flex flex-col z-10">
                 <div className="flex items-center gap-2">
                   <span className="text-yellow-400 text-5xl">👋</span>
                   <span className="text-yellow-400 font-extrabold text-4xl tracking-tighter">clubhouse</span>
                 </div>
                 <span className="text-white text-[10px] ml-14">http://adhu.space/maadhaa27</span>
               </div>

               {/* Center Maadhaa 27 Logo & Live Badge */}
               <div className="absolute left-1/2 -translate-x-1/2 flex items-center z-20">
                  <div className="w-32 h-32 rounded-full bg-[#092B14] border-[4px] border-white flex flex-col items-center justify-center relative shadow-[0_0_0_4px_#092B14]">
                     <span className="text-white font-extrabold text-3xl mb-1" dir="rtl">މާއްދާ</span>
                     <span className="text-white text-[10px] tracking-widest font-bold">MAADHAA</span>
                     <div className="bg-red-600 text-white font-black text-2xl px-5 py-0.5 rounded-full absolute -bottom-4 border-4 border-[#092B14] shadow-md">
                       -27-
                     </div>
                  </div>
                  {/* LIVE Badge */}
                  <div className="bg-red-600 text-white font-black text-[42px] leading-none px-6 py-2 rounded-r-[20px] -ml-8 pl-12 shadow-md flex items-center h-20">
                    LIVE
                  </div>
               </div>

               {/* Right adhu space & Maadhaa27 Text */}
               <div className="z-10 flex flex-col items-end pt-2">
                 <div className="text-white text-sm font-light tracking-wider">
                   adhu<span className="text-red-500 font-medium relative">space</span>
                 </div>
                 <div className="text-[#3CB371] text-5xl font-serif tracking-tight mt-[-4px]">
                   Maadhaa<span className="text-red-600 font-bold">27</span>
                 </div>
                 <div className="text-white/80 text-[10px] mt-1 tracking-wide">http://adhu.space/maadhaa27</div>
               </div>
               
               {/* Right lines */}
               <div className="absolute right-0 top-6 w-20 border-t border-b border-yellow-600/50 h-2"></div>
               <div className="absolute right-0 bottom-6 w-20 border-t border-b border-yellow-600/50 h-2"></div>
            </div>
            
            {/* Light Blue Right Tab */}
            <div className="w-[100px] bg-[#00AEEF] flex flex-col items-center justify-center p-3 text-white border-y-2 border-r-2 border-gray-600 relative">
               <div className="w-12 h-12 bg-[#092B14] rounded-full flex flex-col items-center justify-center text-[10px] font-bold mb-3 border-[2px] border-white text-center leading-tight shadow-md">
                 <span dir="rtl">މާއްދާ</span>
                 <span className="bg-red-600 px-1 rounded-full text-[8px] mt-0.5">-27-</span>
               </div>
               <div className="w-16 h-16 bg-white p-1 rounded-sm shadow-md">
                 {/* Placeholder for QR Code */}
                 <img src="https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=http://adhu.space/maadhaa27" alt="QR" className="w-full h-full object-contain" />
               </div>
            </div>
          </div>
        </section>

        {/* SECTION 1: Single Topic Article */}
        <section className="relative w-full">
           <div className="absolute -left-48 top-1/2 -translate-y-1/2 text-red-600 text-[10px] w-40 text-right font-bold leading-tight z-0">
             Note:<br/>
             Clubhouse - Live<br/>
             Topic Article (Report or News)<br/>
             Topic Article to Publish Daily Here...........{'>'}
          </div>
          <div className="flex w-full h-[400px] shadow-sm">
            <div className="flex-1 bg-white border border-gray-400">
              {/* Content goes here */}
            </div>
            <div className="w-[80px] bg-[#092B14] flex flex-col items-center justify-center rounded-br-[40px] border-b border-r border-t border-gray-400 text-white relative shadow-md">
               <div className="whitespace-nowrap text-3xl font-bold rtl absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 -rotate-90">
                  މާއްދާ 27: ރިޕޯޓް
               </div>
               <div className="absolute bottom-4 right-4 text-yellow-400">
                 <svg className="w-8 h-8" fill="currentColor" viewBox="0 0 20 20"><path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 1.414L10.586 9H7a1 1 0 100 2h3.586l-1.293 1.293a1 1 0 101.414 1.414l3-3a1 1 0 000-1.414z" clipRule="evenodd"></path></svg>
               </div>
            </div>
          </div>
        </section>

        {/* SECTION 2: Latest News Grid (2x2) */}
        <section className="relative w-full">
           <div className="absolute -left-48 top-1/2 -translate-y-1/2 text-red-600 text-[10px] w-40 text-right font-bold leading-tight z-0">
             Note:<br/>
             Latest News / Reports (Politics)<br/>
             Developing news / reports - from local sources to rewrite<br/>
             Topic 4 to Publish Daily Here...........{'>'}
          </div>
          <div className="flex w-full h-[350px] shadow-sm">
            <div className="flex-1 grid grid-cols-2 grid-rows-2 bg-white border-2 border-[#00AEEF]">
               <div className="border-r border-b border-[#00AEEF] p-4"></div>
               <div className="border-b border-[#00AEEF] p-4"></div>
               <div className="border-r border-[#00AEEF] p-4"></div>
               <div className="p-4"></div>
            </div>
            <div className="w-[80px] bg-[#00AEEF] flex flex-col items-center justify-center rounded-br-[40px] text-white relative shadow-md">
               <div className="whitespace-nowrap text-3xl font-bold rtl absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 -rotate-90">
                  އެންމެ ފަހުގެ ޚަބަރު
               </div>
               <div className="absolute bottom-4 right-4 text-[#000B1E]">
                 <svg className="w-8 h-8 bg-white rounded-full p-0.5" fill="currentColor" viewBox="0 0 20 20"><path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 1.414L10.586 9H7a1 1 0 100 2h3.586l-1.293 1.293a1 1 0 101.414 1.414l3-3a1 1 0 000-1.414z" clipRule="evenodd"></path></svg>
               </div>
            </div>
          </div>
        </section>

        {/* SECTION 3: Archives Grid (4x2) */}
        <section className="relative w-full">
           <div className="absolute -left-48 top-1/2 -translate-y-1/2 text-red-600 text-[10px] w-40 text-right font-bold leading-tight z-0">
             Note:<br/>
             Clubhouse Archives (latest 8nEpisodes)<br/>
             To listen Here...........{'>'}
          </div>
          <div className="flex w-full h-[300px] shadow-sm">
            <div className="flex-1 grid grid-cols-4 grid-rows-2 bg-white border-2 border-[#F4B400]">
               {[...Array(8)].map((_, i) => (
                  <div key={i} className={`border-[#F4B400] p-4 ${i < 4 ? 'border-b' : ''} ${i % 4 !== 3 ? 'border-r' : ''}`}>
                    {/* Archive Item Placeholder */}
                  </div>
               ))}
            </div>
            <div className="w-[80px] bg-[#F4B400] flex flex-col items-center justify-center rounded-br-[40px] text-[#092B14] relative shadow-md">
               <div className="whitespace-nowrap text-3xl font-bold rtl absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 -rotate-90 text-center leading-tight">
                  ކުރީގެ އެޕިސޯޑް
               </div>
               <div className="absolute bottom-4 right-4 text-white bg-[#00AEEF] rounded-full">
                 <svg className="w-8 h-8" fill="currentColor" viewBox="0 0 20 20"><path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 1.414L10.586 9H7a1 1 0 100 2h3.586l-1.293 1.293a1 1 0 101.414 1.414l3-3a1 1 0 000-1.414z" clipRule="evenodd"></path></svg>
               </div>
            </div>
          </div>
        </section>

      </main>
    </div>
  );
}
