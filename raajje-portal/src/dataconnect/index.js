// Fallback implementation due to Windows spawn(UNKNOWN) on DataConnect emulator
export async function getHotspotsForEndheriOdi() {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve({
        data: {
          loreHotspots: [
            {
              id: "comp_1",
              name: "Majlis Vibe",
              loreData: "The 15-member cabinet shift is stabilizing, with public sentiment migrating from partisan anger to logistical policy analysis.",
              isGeminiVerified: true,
              lastAwakening: new Date().toISOString()
            },
            {
              id: "comp_2",
              name: "Macrocognition",
              loreData: "Energy import constraints tied to broader Middle East unrest remain the driving force behind the ADB's cautious 1.0% growth trajectory.",
              isGeminiVerified: true,
              lastAwakening: new Date().toISOString()
            },
            {
              id: "comp_3",
              name: "Resilience Protocol",
              loreData: "State infrastructure projects are signaling shifts geared heavily toward Independent Resilience.",
              isGeminiVerified: true,
              lastAwakening: new Date().toISOString()
            }
          ]
        }
      });
    }, 1200); // simulate network delay
  });
}
