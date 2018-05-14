function Queue() {
    this._storage = {};
    this._oldestIndex = 1;
    this._newestIndex = 1;
}

Queue.prototype.size = function() {
    return this._newestIndex - this._oldestIndex;
};

Queue.prototype.enqueue = function(data) {
    this._storage[this._newestIndex] = data;
    this._newestIndex++;
};

Queue.prototype.dequeue = function() {
    let oldestIndex = this._oldestIndex;
    let deletedData = this._storage[oldestIndex];

    delete this._storage[oldestIndex];
    this._oldestIndex++;

    return deletedData;
};

Queue.prototype.getData = function (index) {
    return this._storage[index + this._oldestIndex];
};