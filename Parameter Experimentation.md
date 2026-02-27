Temperature = 0.2 (Low)

Behavior: The model is deterministic, sticks closely to common facts, and avoids creative or unusual outputs.

Example Output:

Machine Learning (ML): Algorithms that learn patterns from data to make predictions.

Deep Learning (DL): Neural networks with multiple layers for tasks like image recognition.

Natural Language Processing (NLP): Computers understanding and generating human language.

Computer Vision: AI analyzing images and videos.

Reinforcement Learning (RL): Training agents through rewards and penalties.

Observation:

Very safe and factual.

Follows standard terminology.

Minimal variation between runs.

Temperature = 0.9 (High)

Behavior: The model is more creative and varied, can introduce slightly unexpected phrasing, examples, or ordering.

Example Output:

Machine Learning (ML): Techniques that allow computers to discover patterns and improve over time.

Deep Learning (DL): Advanced neural networks capable of recognizing images, speech, and even generating art.

Natural Language Processing (NLP): Systems that understand, interpret, and respond to human language.

Computer Vision: AI that interprets visual data, from facial recognition to autonomous cars.

Reinforcement Learning (RL): Teaching AI to make decisions through trial-and-error with rewards.

Observation:

Still factual, but slightly more descriptive and creative.

Uses richer language, additional examples, or broader phrasing.

Outputs can vary more between runs.

Repeat with different top-p values (e.g., 0.5 vs 1).

Top-p = 0.5 (Low)

Behavior: The model only considers the most probable words that together make up 50% of the probability mass.

Result is conservative, focused, and less diverse.

Example Output:

Machine Learning (ML): Algorithms that learn patterns from data.

Deep Learning (DL): Neural networks for tasks like image recognition.

Natural Language Processing (NLP): Computers understanding human language.

Computer Vision: AI interpreting images and videos.

Reinforcement Learning (RL): Training agents with rewards.

Observation:

Very safe and standard.

Little variation between runs.

Avoids unusual phrasing.

Top-p = 1 (High)

Behavior: The model considers all possible words (no probability cutoff).

Result is more varied, creative, and sometimes verbose.

Example Output:

Machine Learning (ML): Techniques where computers learn from data to predict or classify information.

Deep Learning (DL): Advanced neural networks capable of recognizing images, speech, and even generating content.

Natural Language Processing (NLP): AI systems that understand, interpret, and respond to human language.

Computer Vision: Methods for analyzing visual data from images, videos, and real-world environments.

Reinforcement Learning (RL): Training AI agents to make decisions through trial-and-error and rewards.

Observation:

Slightly richer language and more descriptive phrasing.

Outputs may include extra context or examples.

More variability between runs than low top-p.

Record how the style, randomness, and focus of responses change.

Low temperature (0.2) → simple, factual, and consistent style.

High temperature (0.9) → more descriptive, creative, and sometimes conversational.

Low temperature → very low randomness, deterministic outputs.

High temperature → high randomness, outputs vary between runs.

Low top-p (0.5) → conservative style, only common/high-probability words used.

High top-p (1) → more flexible style, includes varied word choices and richer phrasing.

Lower temperature or top-p → highly focused on key points, concise output.

Higher temperature or top-p → slightly less focused, may include extra examples or context.
