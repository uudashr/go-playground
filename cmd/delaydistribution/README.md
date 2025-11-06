# Distribution Math

## 1. Fixed Delay

A fixed delay always adds the same latency value.

### **Equation**
$$
X = c
$$
Where:
- \( X \) = delay (milliseconds)  
- \( c \) = constant delay (e.g., 100 ms)

### **Properties**
$$
\begin{aligned}
\text{Mean}(X) &= c \\
\text{Median}(X) &= c \\
\text{Var}(X) &= 0
\end{aligned}
$$

---

## 2. Uniform Delay

The uniform delay randomly selects a delay between two bounds.

### **Distribution**
$$
X \sim \text{Uniform}(a, b)
$$
Where:
- \( a \) = minimum delay (ms)  
- \( b \) = maximum delay (ms)

### **Probability Density Function**
$$
f_X(x) =
\begin{cases}
\frac{1}{b - a}, & a \le x \le b \\
0, & \text{otherwise}
\end{cases}
$$

### **Sampling**
$$
X = a + (b - a)U
$$
where \( U \sim \text{Uniform}(0,1) \)

### **Properties**
$$
\begin{aligned}
\mathbb{E}[X] &= \frac{a + b}{2} \\
\text{Var}(X) &= \frac{(b - a)^2}{12}
\end{aligned}
$$

---

## 3. Lognormal Delay

A lognormal delay produces natural-looking latency with a long tail.

### **Distribution**
$$
X \sim \text{LogNormal}(\mu, \sigma^2)
$$

### **Parameterization (using median)**
WireMock uses `median` \( m \) and `sigma` \( \sigma \).  
The relationship is:
$$
\mu = \ln(m)
$$

### **Probability Density Function**
$$
f_X(x) = \frac{1}{x\,\sigma\sqrt{2\pi}} \exp\!\left(
-\frac{(\ln x - \mu)^2}{2\sigma^2}
\right), \quad x > 0
$$

### **Sampling**
$$
X = \exp(\mu + \sigma Z)
$$
where \( Z \sim \mathcal{N}(0,1) \)

### **Properties**
$$
\begin{aligned}
\text{Median}(X) &= e^{\mu} = m \\
\mathbb{E}[X] &= e^{\mu + \frac{\sigma^2}{2}} = m\,e^{\frac{\sigma^2}{2}} \\
\text{Var}(X) &= (e^{\sigma^2} - 1)e^{2\mu + \sigma^2}
\end{aligned}
$$

---

## 📈 Percentile Function

To compute the *p-th percentile*:

$$
Q(p) = e^{\mu + \sigma \Phi^{-1}(p)} = m \cdot e^{\sigma \Phi^{-1}(p)}
$$

Where \( \Phi^{-1}(p) \) is the inverse standard normal CDF.

**Example:**  
For \( m = 100 \text{ ms}, \sigma = 0.4 \):

$$
Q(0.95) = 100 \cdot e^{0.4 \cdot 1.645} \approx 184 \text{ ms}
$$
