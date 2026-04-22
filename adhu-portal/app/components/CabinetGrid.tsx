'use client';

const CABINET_MEMBERS = [
  { id: 1, tha: "އިޤްތިޞާދީ ތަރައްޤީ", eng: "Economic Development" },
  { id: 2, tha: "ޚާރިޖީ ކަންތައްތަކާ ބެހޭ", eng: "Foreign Affairs" },
  { id: 3, tha: "ދާޚިލީ ސަލާމަތާއި ފަންނިއްޔާތު", eng: "Homeland Security" },
  { id: 4, tha: "މާލީ ކަންތައްތަކާ ބެހޭ", eng: "Finance" },
  { id: 5, tha: "ތަޢުލީމާ ބެހޭ", eng: "Education" },
  { id: 6, tha: "ޞިއްޙަތާ ބެހޭ", eng: "Health" },
  { id: 7, tha: "އިސްލާމީ ކަންތައްތަކާ ބެހޭ", eng: "Islamic Affairs" },
  { id: 8, tha: "އިމާރާތްކުރުމާއި ބިނާރާ ބެހޭ", eng: "Construction & Infrastructure" },
  { id: 9, tha: "ދަނޑުވެރިކަމާއި ފިޝަރީޒް", eng: "Agriculture & Fisheries" },
  { id: 10, tha: "ތިމާވެއްޓާއި ހަކަތަ", eng: "Environment & Energy" },
  { id: 11, tha: "ޓޫރިޒަމާ ބެހޭ", eng: "Tourism" },
  { id: 12, tha: "ޒުވާނުންނާއި ކުޅިވަރު", eng: "Youth & Sports" },
  { id: 13, tha: "އާއިލީ ތަރައްޤީ", eng: "Social & Family Development" },
  { id: 14, tha: "ގްރާސްރޫޓް އިނީޝިއޭޓިވްސް", eng: "Grassroots Initiatives" },
  { id: 15, tha: "ޓްރާންސްޕޯޓާއި ސިވިލް އޭވިއޭޝަން", eng: "Transport & Civil Aviation" },
];

export default function CabinetGrid({ members }: { members?: any[] }) {
  const displayMembers = members || [
    { id: 1, thaName: "އިޤްތިޞާދީ ތަރައްޤީ", engName: "Economic Development" },
    // ... default mock if none provided
  ];

  return (
    <div className="grid grid-cols-3 gap-2 mt-4">
      {displayMembers.map(member => (
        <div key={member.id} className="glass-card p-2 text-center hover:bg-cyan-900/40 cursor-pointer">
          <div className="kinetic-thaana text-[0.6rem] font-bold whitespace-nowrap overflow-hidden text-ellipsis">
            {member.thaName}
          </div>
          <div className="text-[0.45rem] opacity-60 uppercase tracking-tighter">
            {member.engName}
          </div>
        </div>
      ))}
    </div>
  );
}
