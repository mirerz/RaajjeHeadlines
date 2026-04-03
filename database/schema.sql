-- 1. Table for tracking the various news outlets (Mihaaru, Sun, etc.)
CREATE TABLE IF NOT EXISTS sources (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    base_url TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. The core Article table
CREATE TABLE IF NOT EXISTS articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_id INTEGER REFERENCES sources(id),
    original_url TEXT UNIQUE NOT NULL,
    
    -- Raw content for reference
    raw_headline TEXT NOT NULL,
    raw_body TEXT NOT NULL,
    
    -- AI Generated content (Dhivehi & English)
    rephrased_headline_dv TEXT,
    rephrased_body_dv TEXT,
    summary_en TEXT,
    
    -- Status tracking for Human-in-the-Loop (HITL)
    status VARCHAR(20) DEFAULT 'pending_review' 
        CHECK (status IN ('pending_review', 'approved', 'rejected', 'published')),
    
    -- Editorial metadata
    approved_by_user_id UUID, -- Links to your staff ID
    approved_at TIMESTAMP WITH TIME ZONE,
    
    -- App features
    category VARCHAR(50), -- Politics, Sports, etc.
    is_breaking BOOLEAN DEFAULT FALSE,
    view_count INTEGER DEFAULT 0,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. Indexing for high-performance retrieval (30-min cycles)
CREATE INDEX IF NOT EXISTS idx_articles_status ON articles(status);
CREATE INDEX IF NOT EXISTS idx_articles_created_at ON articles(created_at DESC);
