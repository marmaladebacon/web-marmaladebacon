/* Definitions

### Bias ###
An intercept or offset from an origin. Bias (also known as the bias term) is referred to as b or w0 in machine learning models.
For example, bias is the b in the following formula: y' = b + w1x1 + w2x2 + ...wnxn

### Inference ###
In machine learning, often refers to the process of making predictions by applying the trained model to unlabeled examples. In statistics, inference refers to the process of fitting the parameters of a distribution conditioned on some observed data. (https://en.wikipedia.org/wiki/Statistical_inference)

### Linear Regression ###
A type of regression model that outputs a continuous value from a linear combination of input features.

### Linear Combination ###
In mathematics, a linear combination is an expression constructed from a set of terms by multiplying each term by a constant and adding the results
E.g. a linear combination of x and y would be any expression of the form ax + by, where a and b are constants.

### Weight ###
A coefficient for a feature in a linear model, or an edge in a deep network. The goal of training a linear model is to determine the ideal weight for each feature. If a weight is 0, then its corresponding feature does not contribute to the model.
*/


let x_vals = [];
let y_vals = [];

let m;
let b;

const learningRate = 0.5;

/* ### Stochastic Gradient Descent
  Gradient Descent: A technique to minimize loss by computing the gradients of loss with respect to the model's parameters, conditioned on training data. Informally, gradient descent iteratively adjusts parameters, gradually finding the best combination of weights and bias to minimize loss.

  Stochastic Gradient Descent: A gradient descent algorithm that uses mini-batches. In other words, mini-batch SGD estimates the gradient based on a small subset of the training data. Vanilla SGD uses a mini-batch of size 1.

  How does tensorflow calculate the gradient? Through automatic differentiation: https://stats.stackexchange.com/questions/257746/how-does-tensorflow-tf-train-optimizer-compute-gradients
*/
const optimizer = tf.train.sgd(learningRate);

function setup() {
  createCanvas(400, 400);
  m = tf.variable(tf.scalar(random(1)));
  b = tf.variable(tf.scalar(random(1)));
}

function loss(pred, labels) {
  return pred.sub(labels).square().mean();
}

// x is an array of values
function predict(x) {
  const xs = tf.tensor1d(x);
  //y = mx + b
  const ys = xs.mul(m).add(b);
  return ys;
}

function mousePressed() {
  let x = map(mouseX, 0, width, 0, 1);
  let y = map(mouseY, 0, height, 1, 0);
  x_vals.push(x);
  y_vals.push(y);
}


function draw() {
  /*** Training ************/
  tf.tidy(() => {
    if (x_vals.length > 0){
      const ys = tf.tensor1d(y_vals);
      optimizer.minimize(()=>loss(predict(x_vals), ys));
    }
  });
  /*************************/

  background(0);

  stroke(255);
  strokeWeight(8);
  for (let i = 0; i < x_vals.length; i++) {
    let px = map(x_vals[i], 0, 1, 0, width);
    let py = map(y_vals[i], 0, 1, height, 0);
    point(px, py);
  }

  /*** Prediction ************/
  const lineX = [0,1];

  const ys = tf.tidy(() => predict(lineX));
  let lineY = ys.dataSync();
  ys.dispose();

  /***************************/

  /*** Mapping to canvas coordinates */
  //map is a P5js helper function
  let x1 = map(lineX[0], 0, 1, 0, width);
  let x2 = map(lineX[1], 0, 1, 0, width);

  let y1 = map(lineY[0], 0, 1, height, 0);
  let y2 = map(lineY[1], 0, 1, height, 0);

  strokeWeight(2);
  line(x1,y1,x2,y2);

  //Logging to make sure we are disposing our tensors correctly
  //Should have a constant number of tensors when executing
  //console.log(tf.memory().numTensors);
}